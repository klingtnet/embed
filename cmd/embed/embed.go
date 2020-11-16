package main

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/urfave/cli/v2"
)

func pathToVar(path string) string {
	return fmt.Sprintf("file%x", []byte(path))
}

func encodeFile(data []byte) string {
	return fmt.Sprintf("%q", data)
}

var (
	fileTemplate = template.Must(template.New("").Funcs(template.FuncMap{"pathToVar": pathToVar, "encode": encodeFile}).Parse(`package {{ .Package }}

import (
	"sort"
)

const (
{{- range $path, $data := .Files }}
	{{ pathToVar $path }} = {{ encode $data }}
{{- end }}
)

// Embedded implements github.com/klingtnet/embed/Embed .
type Embedded struct {
	embedMap map[string]string
}

// Embeds stores the embedded data.
var Embeds = Embedded {
	embedMap: map[string]string{
{{- range $path, $_ := .Files }}
		"{{ $path }}": {{ pathToVar $path }},
{{- end }}
	},
}

// Files implements github.com/klingtnet/embed/Embed .
func (e Embedded) Files() []string {
	var fs []string
	for f := range e.embedMap {
		fs = append(fs,f)
	}
	sort.Strings(fs)
	return fs
}

// File implements github.com/klingtnet/embed/Embed .
func (e Embedded) File(path string) []byte {
	file, ok := e.embedMap[path]
	if !ok {
		return nil
	}
	return []byte(file)
}

// FileString implements github.com/klingtnet/embed/Embed .
func (e Embedded) FileString(path string) string {
	return string(e.File(path))
}
`))
)

func readFile(path string) (data []byte, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	data, err = ioutil.ReadAll(f)
	return
}

func embedAction(c *cli.Context) error {
	return embed(c.Context, c.StringSlice("include"), c.String("package"), c.String("destination"))
}

func embed(ctx context.Context, includes []string, packageName, destinationPath string) error {
	files := make(map[string][]byte)

	for _, includePath := range includes {
		info, err := os.Stat(includePath)
		if err != nil {
			return fmt.Errorf("stat: %w", err)
		}
		if info.IsDir() {
			walkFn := func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				data, err := readFile(path)
				if err != nil {
					return fmt.Errorf("readFile: %w", err)
				}
				files[path] = data

				return nil
			}
			err = filepath.Walk(includePath, walkFn)
			if err != nil {
				return fmt.Errorf("filepath.Walk: %w", err)
			}
		} else {
			data, err := readFile(includePath)
			if err != nil {
				return fmt.Errorf("readFile: %w", err)
			}
			files[includePath] = data
		}
	}

	templateData := struct {
		Package string
		Files   map[string][]byte
	}{
		Package: packageName,
		Files:   files,
	}

	buf := bytes.NewBuffer(nil)
	err := fileTemplate.Execute(buf, templateData)
	if err != nil {
		return fmt.Errorf("fileTemplate.Execute: %w", err)
	}
	source, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("format.Source: %w", err)
	}
	dest, err := os.OpenFile(destinationPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("os.OpenFile %q: %w", destinationPath, err)
	}
	defer dest.Close()
	_, err = dest.Write(source)
	if err != nil {
		return fmt.Errorf("dest.Write: %w", err)
	}
	return nil
}

// Version is the build version.
// The actual version is set on build time.
var Version = "unset"

func main() {
	app := cli.App{
		Name:    "embed",
		Version: Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "package",
				Aliases: []string{"p"},
				Usage:   "name of the package the generated Go file is associated to",
				Value:   "main",
			},
			&cli.StringFlag{
				Name:    "destination",
				Aliases: []string{"dest", "d"},
				Usage:   "where to store the generated Go file",
				Value:   "embeds.go",
			},
			&cli.StringSliceFlag{
				Name:     "include",
				Aliases:  []string{"i"},
				Usage:    "paths to embed, directories are stored recursively (can be used multiple times)",
				Required: true,
			},
		},
		Action: embedAction,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
