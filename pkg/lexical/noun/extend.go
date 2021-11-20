package noun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/auxverb/assertor"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/casekind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
)

// Assert asserts the noun.
func (n Noun) Asserted(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	pa := assertor.PoliteAssertor
	if err := conj.Conjunct(pa); err != nil {
		return pa, fmt.Errorf("Noun.Asserted: %v", err)
	}
	return pa, nil
}

// Attributing marks the noun as an attributive.
func (n Noun) Attributing(conj *conjunctor.Conjunctor, c conjunctor.Conjunctable) (extender.Extender, error) {
	if n2, ok := c.(Noun); ok {
		if err := conj.Conjunct(conjugation.NewConjugation(conjugationkind.Attributive), n2); err != nil {
			return nil, fmt.Errorf("Noun.Attributing: %v", err)
		}
		return n2, nil
	}
	return nil, fmt.Errorf("Noun.Attributing: only nouns can be attributed to, got %s", c)
}

// As marks the noun as a case.
func (n Noun) As(conj *conjunctor.Conjunctor, caseKind casekind.CaseKind) (extender.Extender, error) {
	caseMarker, err := particle.From(caseKind)
	if err != nil {
		return nil, fmt.Errorf("Noun.As: %v", err)
	}
	if err := conj.Conjunct(caseMarker); err != nil {
		return nil, fmt.Errorf("Noun.As: %v", err)
	}
	return caseMarker, nil
}
