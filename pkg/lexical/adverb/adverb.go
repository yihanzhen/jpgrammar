package adverb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

// Adverb represents an adverb.
type Adverb struct {
	word.Word
	extender.UnimplementedExtender
}

// NewAdverb creates a new Adverb.
func NewAdverb(writing string) (Adverb, error) {
	w, err := word.NewWord(writing, "")
	if err != nil {
		return Adverb{}, fmt.Errorf("NewAdverb: %v", err)
	}
	n := Adverb{
		Word: w,
		UnimplementedExtender: extender.UnimplementedExtender{
			Name: "Adverb",
		},
	}
	return n, nil
}

// OnWrite implements the Conjunctable interface.
func (adv Adverb) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if !conj.CanStartNewComponent() {
		return nil, fmt.Errorf("Adverb.OnConjunct: sentence can't start new component")
	}
	return &conjunctor.ConjunctorUpdate{
		Inserts:  []conjunctor.Conjunctable{adv},
		WordKind: wordkind.Adverb,
	}, nil
}

// OnWrite implements the Conjunctable interface.
func (adv Adverb) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, adv.Word), nil
}
