package negator

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type NegatorType struct{}

var Negator = NegatorType{}

// OnConjunct implements the Conjunctable interface.
func (n NegatorType) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Adjective && conj.GetConjugationKind() != conjugationkind.Conjunctive {
		return nil, fmt.Errorf("NegatorType.OnConjuct: cannot conjunct noun to wordkind %s and conjugationkind %s", conj.GetWordKind(), conj.GetConjugationKind())
	}
	return &conjunctor.ConjunctorUpdate{
		Inserts:  []conjunctor.Conjunctable{n},
		WordKind: wordkind.AuxVerb,
	}, nil
}

// OnWrite implements the Conjunctable interface.
func (n NegatorType) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, word.MustWord("ない", "ない")), nil
}
