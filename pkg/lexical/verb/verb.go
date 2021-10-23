package verb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Verb struct {
	word.Word
	extender.UnimplementedExtender
	forceTypeOneConj bool
}

type NewVerbOption int

const (
	ForceTypeOneConjugation NewVerbOption = iota
	WriteCanonical
)

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

func (v Verb) OnAppend(c *conjunctor.Conjunctor) error {
	if c.GetWordKind() != wordkind.Particle {
		return fmt.Errorf("Verb.OnAppend: cannot conjunct verb to wordkind: %v", c.GetWordKind())
	}
	c.UpdateWordKind(wordkind.Verb)
	c.Insert(v)
	return nil
}

func (v Verb) OnWrite(_ conjunctor.Conjunctable, words []word.Word) ([]word.Word, error) {
	return append(words, v.Word), nil
}

func (v Verb) Conjugate(ck conjugationkind.ConjugationKind) (word.Word, error) {
	conj := v.getConjugator()
	switch ck {
	case conjugationkind.Imperfective:
		return conj.Imperfective()
	case conjugationkind.Conjunctive:
		return conj.Conjunctive()
	case conjugationkind.Attributive:
		return conj.Attributive()
	case conjugationkind.Terminal:
		return conj.Terminal()
	case conjugationkind.Conditional:
		return conj.Conditional()
	case conjugationkind.Volitional:
		return conj.Volitional()
	}
	return word.Word{}, fmt.Errorf("Verb.Conjugate: conjugationKind not Conjugatable: %v", ck)
}
