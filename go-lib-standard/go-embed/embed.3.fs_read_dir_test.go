package go_embed_test

import (
	"embed"
	"fmt"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed embed-files/*
	embedData embed.FS
)

func TestReadFile(t *testing.T) {
	// Equivalent of bindata.Asset
	content, err := fs.ReadFile(embedData, "embed-files/test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	assert.Equal(t, "List of fruits\nApple\nBanana\nMango\nGrapes\nPineapple\nOrange\nGuava", string(content), "Read file from embed.FS")
}

func TestReadDir(t *testing.T) {
	// Equivalent of bindata.AssetDir
	entries, err := fs.ReadDir(embedData, "embed-files")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		fmt.Printf("Found %s isDir %v\n", entry.Name(), entry.IsDir())
	}

	assert.Equal(t, 2, len(entries), "Read dir from embed.FS")
}
