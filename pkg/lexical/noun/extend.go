package noun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/auxverb"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation/kind"
)

// Assert asserts the noun.
func (n Noun) Asserted(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	pa := auxverb.PoliteAssertor{}
	if err := conj.Conjunct(pa); err != nil {
		return pa, fmt.Errorf("Noun.Asserted: %v", err)
	}
	return pa, nil
}

func (n Noun) Attributing(conj *conjunctor.Conjunctor, c conjunctor.Conjunctable) (extender.Extender, error) {
	if n2, ok := c.(Noun); ok {
		if err := conj.Conjunct(conjugation.NewConjugation(kind.Attributive), n2); err != nil {
			return nil, fmt.Errorf("Noun.Attributing: %v", err)
		}
		return n2, nil
	}

	return nil, fmt.Errorf("Noun.Attributing: only nouns can be attributed to, got %s", c)

}
