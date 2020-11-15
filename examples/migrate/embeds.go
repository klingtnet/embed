package main

import (
	"sort"
)

const (
	file6d6967726174696f6e732f30315f6578616d706c652e75702e73716c = "\x43\x52\x45\x41\x54\x45\x20\x54\x41\x42\x4c\x45\x20\x74\x28\x61\x20\x54\x45\x58\x54\x29\x3b"
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
	return []byte(file)
}

// FileString implements github.com/klingtnet/embed/Embed .
func (e Embedded) FileString(path string) string {
	return string(e.File(path))
}
