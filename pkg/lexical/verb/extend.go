package verb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/auxverb/politer"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation/kind"
)

func (v Verb) Politely(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	if err := conj.Conjunct(conjugation.NewConjugation(kind.Conjunctive), politer.Politer); err != nil {
		return nil, fmt.Errorf("Verb.Politely: %v", err)
	}
	return politer.Politer, nil
}
