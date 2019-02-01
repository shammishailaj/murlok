// +build !wasm

package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilepathsFromDir(t *testing.T) {
	dir := "css-test"
	os.MkdirAll(dir, 0777)
	defer os.RemoveAll(dir)

	h := &Handler{
		webDir: dir,
	}

	assert.Len(t, h.filepathsFromDir(dir, ".css"), 0)

	os.MkdirAll(filepath.Join(dir, "sub"), 0777)
	os.Create(filepath.Join(dir, "test.css"))
	os.Create(filepath.Join(dir, "test.scss"))
	os.Create(filepath.Join(dir, "sub", "sub.css"))

	assert.Contains(t, h.filepathsFromDir(dir, ".css"), "/test.css")
	assert.NotContains(t, h.filepathsFromDir(dir, ".css"), "/test.scss")
	assert.Contains(t, h.filepathsFromDir(dir, ".css"), filepath.Join("/sub", "sub.css"))
}
