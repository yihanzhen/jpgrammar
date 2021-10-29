package auxverb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Politer struct {
	extender.UnimplementedExtender
}

func (p Politer) OnAppend(conj *conjunctor.Conjunctor) error {
	if conj.GetWordKind() != wordkind.Verb && conj.GetConjugationKind() != conjugationkind.Conjunctive {
		return fmt.Errorf("Politer.OnAppend: cannot conjunct Politer to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}
	conj.Insert(p)
	conj.UpdateWordKind(wordkind.AuxVerb)
	conj.UpdateConjugationKind(conjugationkind.Unknown)
	return nil
}

func (p Politer) OnWrite(_ conjunctor.Conjunctable, words []word.Word) ([]word.Word, error) {
	return append(words, word.MustWord("ます", "ます")), nil
}

func (p Politer) Negated(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	np := NegativePoliter{}
	conj.ReplaceHead(np)
	if err := conj.Append(np); err != nil {
		return nil, fmt.Errorf("Politer.Negated: %v", err)
	}
	return np, nil
}

type NegativePoliter struct {
	extender.UnimplementedExtender
}

func (p NegativePoliter) OnAppend(conj *conjunctor.Conjunctor) error {
	if conj.GetWordKind() != wordkind.Verb && conj.GetConjugationKind() != conjugationkind.Conjunctive {
		return fmt.Errorf("Politer.OnAppend: cannot conjunct Politer to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}
	conj.Insert(p)
	conj.UpdateWordKind(wordkind.AuxVerb)
	conj.UpdateConjugationKind(conjugationkind.Unknown)
	return nil
}

func (p NegativePoliter) OnWrite(_ conjunctor.Conjunctable, words []word.Word) ([]word.Word, error) {
	return append(words, word.MustWord("ません", "ません")), nil
}
