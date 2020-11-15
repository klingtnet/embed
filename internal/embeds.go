package internal

import (
	"sort"
)

const (
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30315f746573742e646f776e2e73716c = "\x2d\x2d\x20\x72\x65\x71\x75\x69\x72\x65\x64\x20\x62\x79\x20\x6d\x69\x67\x72\x61\x74\x65\x2f\x64\x72\x69\x76\x65\x72\x5f\x74\x65\x73\x74\x2e\x67\x6f"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30315f746573742e75702e73716c     = "\x2d\x2d\x20\x72\x65\x71\x75\x69\x72\x65\x64\x20\x62\x79\x20\x6d\x69\x67\x72\x61\x74\x65\x2f\x64\x72\x69\x76\x65\x72\x5f\x74\x65\x73\x74\x2e\x67\x6f"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30335f746573742e75702e73716c     = "\x2d\x2d\x20\x72\x65\x71\x75\x69\x72\x65\x64\x20\x62\x79\x20\x6d\x69\x67\x72\x61\x74\x65\x2f\x64\x72\x69\x76\x65\x72\x5f\x74\x65\x73\x74\x2e\x67\x6f"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30345f746573742e646f776e2e73716c = "\x2d\x2d\x20\x72\x65\x71\x75\x69\x72\x65\x64\x20\x62\x79\x20\x6d\x69\x67\x72\x61\x74\x65\x2f\x64\x72\x69\x76\x65\x72\x5f\x74\x65\x73\x74\x2e\x67\x6f"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30345f746573742e75702e73716c     = "\x2d\x2d\x20\x72\x65\x71\x75\x69\x72\x65\x64\x20\x62\x79\x20\x6d\x69\x67\x72\x61\x74\x65\x2f\x64\x72\x69\x76\x65\x72\x5f\x74\x65\x73\x74\x2e\x67\x6f"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30355f746573742e646f776e2e73716c = "\x2d\x2d\x20\x72\x65\x71\x75\x69\x72\x65\x64\x20\x62\x79\x20\x6d\x69\x67\x72\x61\x74\x65\x2f\x64\x72\x69\x76\x65\x72\x5f\x74\x65\x73\x74\x2e\x67\x6f"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30375f746573742e646f776e2e73716c = "\x2d\x2d\x20\x72\x65\x71\x75\x69\x72\x65\x64\x20\x62\x79\x20\x6d\x69\x67\x72\x61\x74\x65\x2f\x64\x72\x69\x76\x65\x72\x5f\x74\x65\x73\x74\x2e\x67\x6f"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30375f746573742e75702e73716c     = "\x2d\x2d\x20\x72\x65\x71\x75\x69\x72\x65\x64\x20\x62\x79\x20\x6d\x69\x67\x72\x61\x74\x65\x2f\x64\x72\x69\x76\x65\x72\x5f\x74\x65\x73\x74\x2e\x67\x6f"
	file696e7465726e616c2f74657374646174612f72616e646f6d2e626c6f62                                 = "\xdc\x24\x18\xc8\xa0\x37\x88\x4e\x25\x9b\x8e\xac\x0f\x7a\xd3\x3d\x12\x07\xc3\x40\x3d\xef\xd1\x75\xcc\xb9\xa5\x54\xa6\xe8\x21\x22\x20\x06\xcd\x14\x77\x65\x23\x83\x40\xe6\x00\x08\x1b\x5f\x8e\x0f\x33\xc8\x6e\xb0\x72\x9d\x26\xdf\xde\xed\xd3\x84\x30\x0f\x20\x6f\x3b\x51\xd9\x46\x9b\xad\x53\xb5\xc8\x38\x39\x20\x9a\xec\x2e\x6f\x6f\xb3\x79\x44\x5e\xc4\xdd\x47\x52\x2c\xd0\x75\x6b\x44\xbc\xf9\x55\x1c\x89\xad\x5b\x50\x65\x84\x04\x13\xb8\xd1\x06\xf0\x1c\x1b\xd5\x42\x53\xda\x3b\xa2\x4f\x8a\x55\xfd\x1c\x11\x09\xd6\xe1\x21"
	file696e7465726e616c2f74657374646174612f746573742e6a736f6e                                     = "\x7b\x22\x6d\x65\x73\x73\x61\x67\x65\x22\x3a\x20\x22\x48\x65\x6c\x6c\x6f\x2c\x20\x57\x6f\x72\x6c\x64\x21\x22\x7d\x0a"
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
