package assertor

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/auxverb/politer"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/verb"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type PoliteAssertorType struct {
	extender.UnimplementedExtender
}

var PoliteAssertor = PoliteAssertorType{
	UnimplementedExtender: extender.NewUnimplementedExtender("polite assertor"),
}

func (p PoliteAssertorType) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Noun && conj.GetWordKind() != wordkind.AdjNoun && conj.GetWordKind() != wordkind.Adjective {
		return nil, fmt.Errorf("PoliteAssertor.OnAppend: cannot conjunct PoliteAssertor to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}

	if conj.GetConjugationKind() != conjugationkind.Unknown {
		return nil, fmt.Errorf("PoliteAssertor.OnAppend: cannot conjunct PoliteAssertor to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}

	return &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AuxVerb,
		ConjugationKind: conjugationkind.Unknown,
		Inserts:         []conjunctor.Conjunctable{p},
	}, nil
}

func (p PoliteAssertorType) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, word.MustWord("です", "です")), nil
}

func (p PoliteAssertorType) Negated(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	v, err := verb.NewVerb("ある", "ある")
	if err != nil {
		return nil, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}

	if err := conj.Update(&conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.Particle,
		ConjugationKind: conjugationkind.Unknown,
		ReplacePrev:     true,
		Inserts:         []conjunctor.Conjunctable{particle.State},
	}); err != nil {
		return nil, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}

	if err := conj.Conjunct(particle.Topic, v, conjugation.NewConjugation(conjugationkind.Conjunctive), politer.Politer); err != nil {
		return nil, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}

	ex, err := politer.Politer.Negated(conj)
	if err != nil {
		return nil, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}
	return ex, nil
}

func (p PoliteAssertorType) Volitionally(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	if err := conj.Update(&conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AuxVerb,
		ConjugationKind: conjugationkind.Volitional,
		ReplacePrev:     true,
		Inserts:         []conjunctor.Conjunctable{VolitionalAssertor},
	}); err != nil {
		return nil, fmt.Errorf("PoliteAssertor.Volitionally: %v", err)
	}
	return VolitionalAssertor, nil
}

type VolitionalAssertorType struct {
	extender.UnimplementedExtender
}

var VolitionalAssertor = VolitionalAssertorType{
	UnimplementedExtender: extender.NewUnimplementedExtender("volitional assertor"),
}

func (v VolitionalAssertorType) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Noun && conj.GetConjugationKind() != conjugationkind.Unknown {
		return nil, fmt.Errorf("PoliteAssertor.OnAppend: cannot conjunct PoliteAssertor to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}

	return &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AuxVerb,
		ConjugationKind: conjugationkind.Volitional,
		Inserts:         []conjunctor.Conjunctable{v},
	}, nil
}

func (p VolitionalAssertorType) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, word.MustWord("でしょう", "でしょう")), nil
}
