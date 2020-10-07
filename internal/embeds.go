package internal

import (
	"encoding/base64"
	"sort"
)

const (
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30315f746573742e646f776e2e73716c = "LS0gcmVxdWlyZWQgYnkgbWlncmF0ZS9kcml2ZXJfdGVzdC5nbw"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30315f746573742e75702e73716c     = "LS0gcmVxdWlyZWQgYnkgbWlncmF0ZS9kcml2ZXJfdGVzdC5nbw"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30335f746573742e75702e73716c     = "LS0gcmVxdWlyZWQgYnkgbWlncmF0ZS9kcml2ZXJfdGVzdC5nbw"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30345f746573742e646f776e2e73716c = "LS0gcmVxdWlyZWQgYnkgbWlncmF0ZS9kcml2ZXJfdGVzdC5nbw"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30345f746573742e75702e73716c     = "LS0gcmVxdWlyZWQgYnkgbWlncmF0ZS9kcml2ZXJfdGVzdC5nbw"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30355f746573742e646f776e2e73716c = "LS0gcmVxdWlyZWQgYnkgbWlncmF0ZS9kcml2ZXJfdGVzdC5nbw"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30375f746573742e646f776e2e73716c = "LS0gcmVxdWlyZWQgYnkgbWlncmF0ZS9kcml2ZXJfdGVzdC5nbw"
	file696e7465726e616c2f74657374646174612f6d6967726174696f6e732f30375f746573742e75702e73716c     = "LS0gcmVxdWlyZWQgYnkgbWlncmF0ZS9kcml2ZXJfdGVzdC5nbw"
	file696e7465726e616c2f74657374646174612f72616e646f6d2e626c6f62                                 = "3CQYyKA3iE4lm46sD3rTPRIHw0A979F1zLmlVKboISIgBs0Ud2Ujg0DmAAgbX44PM8husHKdJt/e7dOEMA8gbztR2UabrVO1yDg5IJrsLm9vs3lEXsTdR1Is0HVrRLz5VRyJrVtQZYQEE7jRBvAcG9VCU9o7ok+KVf0cEQnW4SE"
	file696e7465726e616c2f74657374646174612f746573742e6a736f6e                                     = "eyJtZXNzYWdlIjogIkhlbGxvLCBXb3JsZCEifQo"
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
