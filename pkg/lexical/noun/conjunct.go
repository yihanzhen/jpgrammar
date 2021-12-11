package noun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

// OnConjunct implements the Conjunctable interface.
func (n Noun) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {

	cu := &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.Noun,
		ConjugationKind: conjugationkind.Unknown,
		Inserts:         []conjunctor.Conjunctable{n},
	}

	if conj.NewComponentOK() {
		return cu, nil
	}

	// Compound words.
	if conj.GetWordKind() == wordkind.Noun && conj.GetConjugationKind() == conjugationkind.Unknown {
		return cu, nil
	}

	if conj.GetWordKind() == wordkind.Verb && conj.GetConjugationKind() == conjugationkind.Conjunctive {
		return cu, nil
	}

	// Attributive.
	if conj.GetConjugationKind() == conjugationkind.Attributive {
		return cu, nil
	}

	return nil, fmt.Errorf("Noun(%q).OnAppend: cannot conjunct noun to wordkind %s and conjugationkind %s", n.Write(), conj.GetWordKind(), conj.GetConjugationKind())
}

// OnWrite implements the Conjunctable interface.
func (n Noun) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, n.Word), nil
}
