package adjnoun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

// Noun represents a noun.
type AdjNoun struct {
	word.Word
	extender.UnimplementedExtender
}

// NewAdjNoun creates a new AdjNoun.
func NewAdjNoun(writing string) (AdjNoun, error) {
	w, err := word.NewWord(writing, "")
	if err != nil {
		return AdjNoun{}, fmt.Errorf("NewAdjNoun: %v", err)
	}
	n := AdjNoun{
		Word: w,
		UnimplementedExtender: extender.UnimplementedExtender{
			Name: "AdjNoun",
		},
	}
	return n, nil
}
