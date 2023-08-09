package main

import (
	"log"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilePathAbs(t *testing.T) {
	t.Run("get filepath from file", func(t *testing.T) {
		file_name := "default.md"
		dir, err := filepath.Abs(file_name)
		if err != nil {
			log.Println(err)
		} else {
			dirs := strings.Split(dir, "/")
			assert.Equal(t, file_name, dirs[len(dirs)-1], "filepath Abs returns")
			assert.True(t, len(dirs) > 1, "filepath Abs returns")
		}
	})
}

func TestFilePathBase(t *testing.T) {
	t.Run("get filepath from file", func(t *testing.T) {
		file_name := "default.md"
		dir, err := filepath.Abs(file_name)
		if err != nil {
			log.Println(err)
		} else {
			base := filepath.Base(dir)
			assert.Equal(t, file_name, base, "filepath Base returns")
		}
	})
}

func TestFilePathDir(t *testing.T) {
	t.Run("get filepath from file", func(t *testing.T) {
		file_name := "drafts/default.md"
		path, err := filepath.Abs(file_name)
		if err != nil {
			log.Println(err)
		} else {
			dir := filepath.Dir(path)
			dirs := strings.Split(dir, "/")
			assert.Equal(t, "drafts", dirs[len(dirs)-1], "filepath Dir returns")
		}
	})
}
