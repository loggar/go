package go_embed_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed embed-files/test.txt
	Text_File []byte
)

func TestEmbedSimple(t *testing.T) {
	fmt.Println(string(Text_File))

	assert.Equal(t, "List of fruits\nApple\nBanana\nMango\nGrapes\nPineapple\nOrange\nGuava", string(Text_File), "embed file")
}
