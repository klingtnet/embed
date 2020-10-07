package main

import (
	"encoding/base64"
	"sort"
)

const (
	file6d6967726174696f6e732f30315f6578616d706c652e75702e73716c = "Q1JFQVRFIFRBQkxFIHQoYSBURVhUKTs"
)

// Embedded implements github.com/klingtnet/embed/Embed .
type Embedded struct {
	embedMap map[string]string
}

// Embeds stores the embedded data.
var Embeds = Embedded{
	embedMap: map[string]string{
		"migrations/01_example.up.sql": file6d6967726174696f6e732f30315f6578616d706c652e75702e73716c,
	},
}

// Files implements github.com/klingtnet/embed/Embed .
func (e Embedded) Files() []string {
	var fs []string
	for f := range e.embedMap {
		fs = append(fs, f)
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
	d, err := base64.RawStdEncoding.DecodeString(file)
	if err != nil {
		panic(err)
	}
	return d
}

// FileString implements github.com/klingtnet/embed/Embed .
func (e Embedded) FileString(path string) string {
	return string(e.File(path))
}
