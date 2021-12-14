package adjective

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

// OnConjunct implements the Conjunctable interface.
func (a Adjective) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	cu := &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.Adjective,
		ConjugationKind: conjugationkind.Unknown,
		Inserts:         []conjunctor.Conjunctable{a},
	}

	if conj.NewComponentOK() {
		return cu, nil
	}

	return nil, fmt.Errorf("AdjNoun(%q).OnAppend: cannot conjunct noun to wordkind %s and conjugationkind %s", a.Write(), conj.GetWordKind(), conj.GetConjugationKind())
}

// OnWrite implements the Conjunctable interface.
func (a Adjective) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, a.Word), nil
}
