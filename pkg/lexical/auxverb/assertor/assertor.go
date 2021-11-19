package assertor

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/auxverb/politer"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation/kind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/verb"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type PoliteAssertor struct {
	extender.UnimplementedExtender
}

func (p PoliteAssertor) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Noun && conj.GetConjugationKind() != kind.Unknown {
		return nil, fmt.Errorf("PoliteAssertor.OnAppend: cannot conjunct PoliteAssertor to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}

	return &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AuxVerb,
		ConjugationKind: kind.Unknown,
		Inserts:         []conjunctor.Conjunctable{p},
	}, nil
}

func (p PoliteAssertor) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, word.MustWord("です", "です")), nil
}

func (p PoliteAssertor) Negated(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	v, err := verb.NewVerb("ある", "ある")
	if err != nil {
		return nil, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}

	if err := conj.Update(&conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.Particle,
		ConjugationKind: kind.Unknown,
		ReplacePrev:     true,
		Inserts:         []conjunctor.Conjunctable{particle.State},
	}); err != nil {
		return nil, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}

	if err := conj.Conjunct(particle.Topic, v, conjugation.NewConjugation(kind.Conjunctive), politer.Politer); err != nil {
		return nil, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}

	ex, err := politer.Politer.Negated(conj)
	if err != nil {
		return nil, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}
	return ex, nil
}
