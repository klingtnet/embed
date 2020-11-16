package internal

import (
	"sort"
)

const (
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30315f746573742e646f776e2e73716c = "-- required by migrate/driver_test.go"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30315f746573742e75702e73716c     = "-- required by migrate/driver_test.go"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30335f746573742e75702e73716c     = "-- required by migrate/driver_test.go"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30345f746573742e646f776e2e73716c = "-- required by migrate/driver_test.go"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30345f746573742e75702e73716c     = "-- required by migrate/driver_test.go"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30355f746573742e646f776e2e73716c = "-- required by migrate/driver_test.go"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30375f746573742e646f776e2e73716c = "-- required by migrate/driver_test.go"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30375f746573742e75702e73716c     = "-- required by migrate/driver_test.go"
	file696e7465726e616c2f74657374646174612f72616e646f6d2e626c6f62                                 = "\xdc$\x18Ƞ7\x88N%\x9b\x8e\xac\x0fz\xd3=\x12\a\xc3@=\xef\xd1u̹\xa5T\xa6\xe8!\" \x06\xcd\x14we#\x83@\xe6\x00\b\x1b_\x8e\x0f3\xc8n\xb0r\x9d&\xdf\xde\xedӄ0\x0f o;Q\xd9F\x9b\xadS\xb5\xc889 \x9a\xec.oo\xb3yD^\xc4\xddGR,\xd0ukD\xbc\xf9U\x1c\x89\xad[Pe\x84\x04\x13\xb8\xd1\x06\xf0\x1c\x1b\xd5BS\xda;\xa2O\x8aU\xfd\x1c\x11\t\xd6\xe1!"
	file696e7465726e616c2f74657374646174612f746573742e6a736f6e                                     = "{\"message\": \"Hello, World!\"}\n"
)

// Embedded implements github.com/klingtnet/embed/Embed .
type Embedded struct {
	embedMap map[string]string
}

// Embeds stores the embedded data.
var Embeds = Embedded{
	embedMap: map[string]string{
		"internal/testdata/migrations/01_test.down.sql": file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30315f746573742e646f776e2e73716c,
		"internal/testdata/migrations/01_test.up.sql":   file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30315f746573742e75702e73716c,
		"internal/testdata/migrations/03_test.up.sql":   file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30335f746573742e75702e73716c,
		"internal/testdata/migrations/04_test.down.sql": file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30345f746573742e646f776e2e73716c,
		"internal/testdata/migrations/04_test.up.sql":   file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30345f746573742e75702e73716c,
		"internal/testdata/migrations/05_test.down.sql": file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30355f746573742e646f776e2e73716c,
		"internal/testdata/migrations/07_test.down.sql": file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30375f746573742e646f776e2e73716c,
		"internal/testdata/migrations/07_test.up.sql":   file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30375f746573742e75702e73716c,
		"internal/testdata/random.blob":                 file696e7465726e616c2f74657374646174612f72616e646f6d2e626c6f62,
		"internal/testdata/test.json":                   file696e7465726e616c2f74657374646174612f746573742e6a736f6e,
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
