package go_embed_test

import (
	"embed"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed embed-files/*
	folder embed.FS
)

func TestEmbedFS(t *testing.T) {
	data1, _ := folder.ReadFile("embed-files/test.txt")
	data2, _ := folder.ReadFile("embed-files/nested_dir/other.txt")

	assert.Equal(t, "List of fruits\nApple\nBanana\nMango\nGrapes\nPineapple\nOrange\nGuava", string(data1), "embed.FS file")
	assert.Equal(t, "other file content", string(data2), "embed.FS file")

	data3, err := folder.ReadFile("embed-files/nested_dir/not_exist.txt")

	assert.Equal(t, "", string(data3), "embed.FS file which does not exist")
	assert.Equal(t, errors.New("open embed-files/nested_dir/not_exist.txt: file does not exist").Error(), err.Error(), "embed.FS file which does not exist")
}
