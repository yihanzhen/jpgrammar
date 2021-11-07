package noun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

// Noun represents a noun.
type Noun struct {
	word.Word
	extender.UnimplementedExtender
}

// NewNoun creates a new Noun.
func NewNoun(writing string) (Noun, error) {
	w, err := word.NewWord(writing, "")
	if err != nil {
		return Noun{}, fmt.Errorf("NewNoun: %v", err)
	}
	n := Noun{
		Word: w,
		UnimplementedExtender: extender.UnimplementedExtender{
			Name: "Noun",
		},
	}
	return n, nil
}

// Omitted represents an omitted noun.
var Omitted Noun = Noun{
	Word: word.Omitted,
	UnimplementedExtender: extender.UnimplementedExtender{
		Name: "Noun",
	},
}
