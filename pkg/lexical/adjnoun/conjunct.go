package adjnoun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

// OnConjunct implements the Conjunctable interface.
func (n AdjNoun) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {

	cu := &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AdjNoun,
		ConjugationKind: conjugationkind.Unknown,
		Inserts:         []conjunctor.Conjunctable{n},
	}

	if conj.NewComponentOK() {
		return cu, nil
	}

	return nil, fmt.Errorf("AdjNoun(%q).OnAppend: cannot conjunct noun to wordkind %s and conjugationkind %s", n.Write(), conj.GetWordKind(), conj.GetConjugationKind())
}

// OnWrite implements the Conjunctable interface.
func (n AdjNoun) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, n.Word), nil
}
