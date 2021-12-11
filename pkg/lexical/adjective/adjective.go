package adjective

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

// Adjective represents a noun.
type Adjective struct {
	word.Word
	extender.UnimplementedExtender
}

// NewAdjective creates a new Noun.
func NewAdjective(writing string) (Adjective, error) {
	w, err := word.NewWord(writing, "")
	if err != nil {
		return Adjective{}, fmt.Errorf("NewAdjective: %v", err)
	}
	if !w.CheckSuffix("い") {
		return Adjective{}, fmt.Errorf("NewAdjective: got %v, want writing end with い", writing)
	}
	n := Adjective{
		Word: w,
		UnimplementedExtender: extender.UnimplementedExtender{
			Name: "Adjective",
		},
	}
	return n, nil
}
