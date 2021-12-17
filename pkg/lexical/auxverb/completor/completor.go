package completor

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/auxverb/assertor"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type CompletorType struct {
	extender.UnimplementedExtender
}

var Completor = CompletorType{
	UnimplementedExtender: extender.NewUnimplementedExtender("completor"),
}

func (c CompletorType) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetConjugationKind() != conjugationkind.Conjunctive {
		return nil, fmt.Errorf("Completor.OnAppend: cannot conjunct Completor to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}

	return &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AuxVerb,
		ConjugationKind: conjugationkind.Unknown,
		Inserts:         []conjunctor.Conjunctable{c},
	}, nil
}

func (c CompletorType) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, word.MustWord("た", "")), nil
}

func (c CompletorType) Politely(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	// TODO: only adjective + た + です is valid. Add support for checking the previous word is an adjective.
	if err := conj.Conjunct(assertor.PoliteAssertor); err != nil {
		return c, fmt.Errorf("Completor.Politely: %v", err)
	}
	return assertor.PoliteAssertor, nil
}
