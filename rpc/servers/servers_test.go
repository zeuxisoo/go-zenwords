package servers

import (
	"testing"

	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
)

func TestMain(m *testing.M) {
	keywords.NewKeywords("../../words.txt")
}
