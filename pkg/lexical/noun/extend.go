package noun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/auxverb"
)

// Assert asserts the noun.
func (n Noun) Asserted(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	pa := auxverb.PoliteAssertor{}
	if err := conj.Conjunct(pa); err != nil {
		return nil, fmt.Errorf("Noun.Asserted: %v", err)
	}
	return pa, nil
}
