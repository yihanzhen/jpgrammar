package adjnoun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/auxverb/assertor"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/noun"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
)

// Assert asserts the noun.
func (a AdjNoun) Asserted(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	pa := assertor.PoliteAssertor
	if err := conj.Conjunct(pa); err != nil {
		return pa, fmt.Errorf("AdjNoun.Asserted: %v", err)
	}
	return pa, nil
}

// Attributing marks the noun as an attributive.
func (a AdjNoun) Attributing(conj *conjunctor.Conjunctor, c conjunctor.Conjunctable) (extender.Extender, error) {
	if n2, ok := c.(noun.Noun); ok {
		if err := conj.Conjunct(conjugation.NewConjugation(conjugationkind.Attributive), n2); err != nil {
			return nil, fmt.Errorf("AdjNoun.Attributing: %v", err)
		}
		return n2, nil
	}
	return nil, fmt.Errorf("AdjNoun.Attributing: only nouns can be attributed to, got %s", c)
}
