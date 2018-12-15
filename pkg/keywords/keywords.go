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

// ExtraSearch return the result string and match state base on the key
func ExtraSearch(key string) (result string, matched bool) {
	index, matched := Builder.ExactMatchSearch(key)

	if matched == false {
		return "", matched
	}

	return Words[index], matched
}

// PrefixSearch return the result set of string and match state base on the key
func PrefixSearch(key string) (results []string, matched bool) {
	indexes := Builder.CommonPrefixSearch(key)
	values  := []string{}

	if len(indexes) <= 0 {
		return values, false
	}

	for _, position := range indexes {
		values = append(values, Words[position])
	}

	return values, true
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
