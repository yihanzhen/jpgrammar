package verb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/casekind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Verb struct {
	word.Word
	extender.UnimplementedExtender
	forceConjTypeOne bool
}

type NewVerbOption int

const (
	ForceTypeOneConjugation NewVerbOption = iota
)

func NewVerb(writing, conjRef string, opts ...NewVerbOption) (Verb, error) {
	var forceTypeOneConj bool
	for _, opt := range opts {
		if opt == ForceTypeOneConjugation {
			forceTypeOneConj = true
		}
	}
	if conjRef == "" {
		return Verb{}, fmt.Errorf("NewVerb: conjRef cannot be empty")
	}
	w, err := word.NewWord(writing, conjRef)
	if err != nil {
		return Verb{}, fmt.Errorf("NewVerb: %v", err)
	}
	if !w.CheckLastRuneCol(2) {
		return Verb{}, fmt.Errorf("NewVerb: last rune of word is not in col 2")
	}
	v := Verb{
		Word:             w,
		forceConjTypeOne: forceTypeOneConj,
	}
	return v, nil
}

func (v Verb) OnConjunct(c *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	cu := &conjunctor.ConjunctorUpdate{
		WordKind: wordkind.Verb,
		Inserts:  []conjunctor.Conjunctable{v},
	}
	if c.GetCaseKind() != casekind.Unknown {
		return cu, nil
	}
	if c.GetWordKind() == wordkind.Particle || c.GetWordKind() == wordkind.Adverb {
		return cu, nil
	}
	return nil, fmt.Errorf("Verb.OnConjunct: cannot conjunct verb to wordkind: %v", c.GetWordKind())
}

func (v Verb) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, v.Word), nil
}
