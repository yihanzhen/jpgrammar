package verb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Verb struct {
	conjunctor.DefaultConjunctable
	word.Word
	forceTypeOneConj bool
}

type NewVerbOption int

const (
	ForceTypeOneConjugation NewVerbOption = iota
	WriteCanonical
)

func (v Verb) GetWordKind() wordkind.WordKind {
	return wordkind.Verb
}

func (v Verb) CheckPrev(c *conjunctor.Conjunctor, prev conjunctor.Conjunctable) error {
	return nil
}

func (v Verb) CheckNext(c *conjunctor.Conjunctor, next conjunctor.Conjunctable) error {
	return nil
}

func (v Verb) OnConjunct(prev, next conjunctor.Conjunctable) ([]conjunctor.Conjunctable, error) {
	return []conjunctor.Conjunctable{prev, v, next}, nil
}

func (v Verb) OnWrite(words []word.Word) []word.Word {
	return append(words, v.Word)
}

func (v Verb) GetExtender() extender.Extender {
	return &VerbExtender{}
}

func NewVerb(canonical, display string, opts ...NewVerbOption) (Verb, error) {
	var forceTypeOneConj bool
	for _, opt := range opts {
		if opt == WriteCanonical {
			display = canonical
		}
		if opt == ForceTypeOneConjugation {
			forceTypeOneConj = true
		}
	}

	w, err := word.NewWord(canonical, display)
	if err != nil {
		return Verb{}, fmt.Errorf("NewVerb: %v", err)
	}
	if !w.CheckLastRuneCol(2) {
		return Verb{}, fmt.Errorf("NewVerb: last rune of word is not in col 2")
	}
	v := Verb{
		Word:             w,
		forceTypeOneConj: forceTypeOneConj,
	}
	return v, nil
}

func (v Verb) GetConjugator() VerbConjugator {
	if v.forceTypeOneConj {
		return &TypeOneVerbConjugator{
			verb: v,
		}
	}
	if !v.CheckLastRune('る') {
		return &TypeOneVerbConjugator{
			verb: v,
		}
	}
	if !v.CheckRune(word.NthLastRune(1), word.IsCol(1)) && !v.CheckRune(word.NthLastRune(1), word.IsCol(3)) {
		return &TypeOneVerbConjugator{
			verb: v,
		}
	}
	return &TypeTwoVerbConjugator{
		verb: v,
	}
}
