package tester

import (
	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
)

// CreateRPC will prepare all related data for RPC testing
func CreateRPC() {
	keywords.NewKeywords("../../words.txt")
}
