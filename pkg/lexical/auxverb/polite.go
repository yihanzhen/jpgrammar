package auxverb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation/kind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Politer struct {
	extender.UnimplementedExtender
}

func (p Politer) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Verb && conj.GetConjugationKind() != kind.Conjunctive {
		return nil, fmt.Errorf("Politer.OnAppend: cannot conjunct Politer to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}

	return &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AuxVerb,
		ConjugationKind: kind.Unknown,
		Inserts:         []conjunctor.Conjunctable{p},
	}, nil
}

func (p Politer) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, word.MustWord("ます", "ます")), nil
}

func (p Politer) Negated(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	np := NegativePoliter{}
	if err := conj.Update(&conjunctor.ConjunctorUpdate{
		ReplacePrev: true,
		Inserts:     []conjunctor.Conjunctable{np},
	}); err != nil {
		return nil, fmt.Errorf("Politer.Negated: %v", err)
	}
	return np, nil
}

type NegativePoliter struct {
	extender.UnimplementedExtender
}

func (p NegativePoliter) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Verb && conj.GetConjugationKind() != kind.Conjunctive {
		return nil, fmt.Errorf("Politer.OnAppend: cannot conjunct Politer to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}
	return &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AuxVerb,
		ConjugationKind: kind.Unknown,
		Inserts:         []conjunctor.Conjunctable{p},
	}, nil
}

func (p NegativePoliter) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, word.MustWord("ません", "ません")), nil
}
