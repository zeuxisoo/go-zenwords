package keywords

import (
	"os"
	"bufio"
	"io"
	"bytes"
	"log"
	"sort"

	"github.com/importcjj/sensitive"
)

var (
	SensitiveFilter *sensitive.Filter
	Words 			[]string
)

// NewKeywords to setup the base keywords environment
func NewKeywords(keywordFile string) {
	words, err := loadFile(keywordFile)
	if err != nil {
		log.Fatalf("Cannot load the keyword file: %v", err)
	}
	sort.Strings(words)

	filter := sensitive.New()
	filter.AddWord(words...)

	SensitiveFilter = filter
	Words = words
}

// Filter will return the filtered result
func Filter(content string) string {
	return SensitiveFilter.Replace(content, rune('*'))
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
