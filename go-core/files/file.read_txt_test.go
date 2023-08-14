package core_files

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadTxtAsString(t *testing.T) {
	text, err := os.ReadFile("_sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	assert.NoError(t, err, "read file")

	assert.Equal(t, "a\nb\nc\n1\n2\n3\n", string(text), "read a text file as a string")
}

func TestReadFileByLine(t *testing.T) {
	f, err := os.Open("_sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	line_list := []string{}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line_list = append(line_list, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	expected := []string{"a", "b", "c", "1", "2", "3"}

	for i, line := range line_list {
		assert.Equal(t, expected[i], line, "read a file line by line")
	}
}

func TestReadFileByWord(t *testing.T) {
	f, err := os.Open("_sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wordlist := []string{}
	for scanner.Scan() {
		word := scanner.Text()
		wordlist = append(wordlist, word)
	}

	expected := []string{"a", "b", "c", "1", "2", "3"}

	assert.Equal(t, expected, wordlist, "read a file word by word")
}

func TestReadFileByChar(t *testing.T) {
	LineDelimiter := "\n"
	if runtime.GOOS == "windows" {
		LineDelimiter = "\r\n"
	}

	f, err := os.Open("_sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	char_list := []byte{}
	for scanner.Scan() {
		char := byte(scanner.Text()[0])
		char_list = append(char_list, char)
	}

	expected_char := []string{"a", LineDelimiter, "b", LineDelimiter, "c", LineDelimiter, "1", LineDelimiter, "2", LineDelimiter, "3", LineDelimiter}
	expected := []uint8{}

	for _, c := range expected_char {
		expected = append(expected, byte(c[0]))
	}

	assert.Equal(t, expected, char_list, "read a file character by character")
}
