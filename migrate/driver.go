package migrate

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4/source"

	"github.com/klingtnet/embed"
)

func init() {
	source.Register("embed", &EmbedDriver{})
}

// EmbedDriver implements migrate/source.Driver .
type EmbedDriver struct {
	embeds        embed.Embed
	migrationPath string
	migrations    *source.Migrations
}

// Open implements migrate/source.Driver .
func (ed *EmbedDriver) Open(url string) (source.Driver, error) {
	return nil, fmt.Errorf("unimplemented, use WithInstance instead")
}

// WithInstance instantiates a new migration driver with the given embed.Embeds .
func WithInstance(instance interface{}) (source.Driver, error) {
	embeds, ok := instance.(embed.Embed)
	if !ok {
		return nil, fmt.Errorf("instance is not of type embed.Embed")
	}

	ms := source.NewMigrations()
	for _, file := range embeds.Files() {
		m, err := source.DefaultParse(filepath.Base(file))
		if err != nil {
			// ignore unparseable files
			continue
		}
		m.Raw = file

		ok := ms.Append(m)
		if !ok {
			return nil, source.ErrDuplicateMigration{
				Migration: *m,
			}
		}
	}

	return &EmbedDriver{
		embeds:     embeds,
		migrations: ms,
	}, nil
}

// Close implements migrate/source.Driver .
func (ed *EmbedDriver) Close() error {
	return nil
}

// First implements migrate/source.Driver .
func (ed *EmbedDriver) First() (version uint, err error) {
	if version, ok := ed.migrations.First(); ok {
		return version, nil
	}
	return 0, os.ErrNotExist
}

// Prev implements migrate/source.Driver .
func (ed *EmbedDriver) Prev(version uint) (prevVersion uint, err error) {
	if version, ok := ed.migrations.Prev(version); ok {
		return version, nil
	}
	return 0, os.ErrNotExist
}

// Next implements migrate/source.Driver .
func (ed *EmbedDriver) Next(version uint) (nextVersion uint, err error) {
	if version, ok := ed.migrations.Next(version); ok {
		return version, nil
	}
	return 0, os.ErrNotExist
}

// ReadUp implements migrate/source.Driver .
func (ed *EmbedDriver) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := ed.migrations.Up(version); ok {
		return ioutil.NopCloser(bytes.NewBuffer(ed.embeds.File(m.Raw))), m.Identifier, nil
	}
	return nil, "", os.ErrNotExist
}

// ReadDown implements migrate/source.Driver .
func (ed *EmbedDriver) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := ed.migrations.Down(version); ok {
		return ioutil.NopCloser(bytes.NewBuffer(ed.embeds.File(m.Raw))), m.Identifier, nil
	}
	return nil, "", os.ErrNotExist
}
