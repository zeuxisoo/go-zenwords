package keywords

import (
	"os"
	"bufio"
	"io"
	"bytes"
	"log"
)

var (
	Words	[][]rune
)

// NewKeywords to setup the base keywords environment
func NewKeywords(keywordFile string) {
	var err error

	Words, err = loadFile(keywordFile)
	if err != nil {
		log.Fatalf("Cannot read the words.txt file: %v\n", err)
	}
}

func loadFile(filePath string) ([][]rune, error) {
	words := [][]rune{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}else{
			line  = bytes.TrimSpace(line)
			words = append(words, bytes.Runes(line))
		}
	}

	return words, nil
}
