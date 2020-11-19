package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func BenchmarkEncoding(b *testing.B) {
	bCases := []struct {
		name string
		N    int64
	}{
		{"100k", 102_400},
		{"1MiB", 1_048_576},
		{"10MiB", 10_485_760},
		{"100MiB", 104_857_600},
		{"1GiB", 1_073_741_824},
	}

	randReader := rand.New(rand.NewSource(time.Now().Unix()))
	for _, bCase := range bCases {
		testData := bytes.NewBuffer(make([]byte, bCase.N))
		_, err := io.CopyN(testData, randReader, bCase.N)
		if err != nil {
			b.Fatal(err)
		}
		b.Run(bCase.name, func(b *testing.B) {
			b.ReportAllocs()
			_ = encodeFile(testData.Bytes())
		})
	}
}

func BenchmarkEmbed(b *testing.B) {
	bCases := []struct {
		name string
		noFiles,
		N int64
	}{
		{"100k", 1000, 102_400},
		{"1MiB", 100, 1_048_576},
		{"10MiB", 10, 10_485_760},
		{"100MiB", 1, 104_857_600},
		// go/format.Source source will panic for a 1GiB file
		// {"1GiB", 1_073_741_824},
	}

	randReader := rand.New(rand.NewSource(time.Now().Unix()))
	for _, bCase := range bCases {
		b.Run(bCase.name, func(b *testing.B) {
			dir := b.TempDir()
			for i := int64(0); i < bCase.noFiles; i++ {
				testfile := filepath.Join(dir, fmt.Sprintf("%s%d.bin", bCase.name, i))
				f, err := os.OpenFile(testfile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
				if err != nil {
					b.Fatalf("creating test file %q failed: %s", f.Name(), err.Error())
				}
				n, err := io.CopyN(f, randReader, bCase.N)
				if err != nil {
					b.Fatalf("writing random test file %q failed: %s", f.Name(), err)
				}
				if n != bCase.N {
					b.Fatalf("expected %d random bytes to be written but was %d", bCase.N, n)
				}
				err = f.Close()
				if err != nil {
					b.Fatalf("closing test file %q failed: %s", f.Name(), err)
				}
			}

			b.ReportAllocs()
			b.ResetTimer()
			err := embed(context.Background(), []string{dir}, "main", filepath.Join(dir, "embeds.go"))
			b.StopTimer()
			if err != nil {
				b.Fatal(err)
			}
		})
	}
}
