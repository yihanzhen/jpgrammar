package adjective

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/auxverb/assertor"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/auxverb/negator"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/noun"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
)

// Negated negates the adjective.
func (adj Adjective) Negated(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	if err := conj.Conjunct(conjugation.NewConjugation(conjugationkind.Conjunctive)); err != nil {
		return nil, fmt.Errorf("Adjective.Negated: %v", err)
	}
	if err := conj.Conjunct(negator.Negator); err != nil {

	}
	return nil, nil
}

// Attributing marks the adjective as an attributive.
func (adj Adjective) Attributing(conj *conjunctor.Conjunctor, c conjunctor.Conjunctable) (extender.Extender, error) {
	if n2, ok := c.(noun.Noun); ok {
		if err := conj.Conjunct(conjugation.NewConjugation(conjugationkind.Attributive), n2); err != nil {
			return nil, fmt.Errorf("Adjective.Attributing: %v", err)
		}
		return n2, nil
	}
	return nil, fmt.Errorf("Adjective.Attributing: only nouns can be attributed to, got %s", c)
}

// Politely makes the tone of the sentence polite.
func (adj Adjective) Politely(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	fmt.Println("here")
	pa := assertor.PoliteAssertor
	if err := conj.Conjunct(pa); err != nil {
		return pa, fmt.Errorf("Adjective.Politely: %v", err)
	}
	return pa, nil
}
