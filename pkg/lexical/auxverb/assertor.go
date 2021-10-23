package auxverb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/verb"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type PoliteAssertor struct {
	extender.UnimplementedExtender
}

func (p PoliteAssertor) OnAppend(conj *conjunctor.Conjunctor) error {
	if conj.GetWordKind() != wordkind.Noun && conj.GetConjugationKind() != conjugationkind.Unknown {
		return fmt.Errorf("PoliteAssertor.OnAppend: cannot conjunct PoliteAssertor to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}
	conj.Insert(p)
	conj.UpdateWordKind(wordkind.AuxVerb)
	conj.UpdateConjugationKind(conjugationkind.Unknown)
	return nil
}

func (p PoliteAssertor) OnWrite(_ conjunctor.Conjunctable, words []word.Word) ([]word.Word, error) {
	return append(words, word.MustWord("です", "です")), nil
}

func (e *PoliteAssertor) Negated(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	v, err := verb.NewVerb("ある", "ある")
	if err != nil {
		return e, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}
	if err := conj.RemoveHead(); err != nil {
		return e, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}
	if err := conj.Append(particle.State, particle.Topic, v, Politer{}); err != nil {
		return e, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}
	ex, err := Politer{}.Negated(conj)
	if err != nil {
		return e, fmt.Errorf("PoliteAssertor.Negated: %v", err)
	}
	return ex, nil
}
