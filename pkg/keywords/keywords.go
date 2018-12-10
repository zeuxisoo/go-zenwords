package keywords

import (
	"os"
	"bufio"
	"io"
	"bytes"
	"log"
	"sort"

	"github.com/euclidr/darts"
)

var (
	Words		[]string
	Builder		darts.DoubleArrayBuilder
)

// NewKeywords to setup the base keywords environment
func NewKeywords(keywordFile string) {
	words, err := loadFile(keywordFile)
	if err != nil {
		log.Fatalf("Cannot read the words.txt file: %v\n", err)
	}
	sort.Strings(words)

	builder := darts.DoubleArrayBuilder{}
	builder.Build(words)

	Words    = words
	Builder  = builder
}

// Search return the result string and match state base on content
func Search(content string) (result string, matched bool) {
	index, matched := Builder.ExactMatchSearch(content)

	if matched == false {
		return "", matched
	}else{
		return Words[index], matched
	}
}

func loadFile(filePath string) ([]string, error) {
	words := []string{}

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
			words = append(words, string(line))
		}
	}

	return words, nil
}
