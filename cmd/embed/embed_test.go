package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"testing"

	"github.com/klingtnet/embed/internal"
	"github.com/stretchr/testify/require"
)

func TestEmbedCommand(t *testing.T) {
	tf, err := ioutil.TempFile("", t.Name()+"-*")
	require.NoError(t, err, "ioutil.TempFile")
	t.Cleanup(func() {
		_ = tf.Close()
		_ = os.Remove(tf.Name())
	})
	err = embed(context.Background(), []string{"../../go.mod"}, "mypackage", tf.Name())
	require.NoError(t, err)
	io.Copy(os.Stdout, tf)
}

func TestEmbeds(t *testing.T) {
	expectedFiles := []string{
		"internal/testdata/test.json",
		"internal/testdata/random.blob",
		"internal/testdata/migrations/01_test.up.sql",
		"internal/testdata/migrations/03_test.up.sql",
		"internal/testdata/migrations/04_test.up.sql",
		"internal/testdata/migrations/04_test.down.sql",
		"internal/testdata/migrations/07_test.up.sql",
		"internal/testdata/migrations/01_test.down.sql",
		"internal/testdata/migrations/05_test.down.sql",
		"internal/testdata/migrations/07_test.down.sql",
	}
	sort.Strings(expectedFiles)
	require.Equal(t, expectedFiles, internal.Embeds.Files())

	testJSON := struct {
		Message string `json:"message"`
	}{}
	err := json.Unmarshal(internal.Embeds.File("internal/testdata/test.json"), &testJSON)
	require.NoError(t, err, "json.Unmarshal")
	require.Equal(t, "Hello, World!", testJSON.Message, "unexpected content of test.json")

	sum := sha256.Sum256(internal.Embeds.File("internal/testdata/random.blob"))
	require.Equal(t, "d3e63e3097f1da36b0a83a9115e23da333db5eac1ff4581f50593d94e368b56b", hex.EncodeToString(sum[:]), "sha256 of random blob differs")
}
