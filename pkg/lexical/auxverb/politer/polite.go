package politer

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type PoliterType struct {
	extender.UnimplementedExtender
}

var Politer = PoliterType{}

func (p PoliterType) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Verb && conj.GetConjugationKind() != conjugationkind.Conjunctive {
		return nil, fmt.Errorf("Politer.OnConjunct: cannot conjunct Politer to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}

	return &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AuxVerb,
		ConjugationKind: conjugationkind.Unknown,
		Inserts:         []conjunctor.Conjunctable{p},
	}, nil
}

func (p PoliterType) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, word.MustWord("ます", "ます")), nil
}

func (p PoliterType) Negated(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	np := NegativePoliter{}
	if err := conj.Update(&conjunctor.ConjunctorUpdate{
		ReplacePrev: true,
		Inserts:     []conjunctor.Conjunctable{np},
	}); err != nil {
		return nil, fmt.Errorf("Politer.Negated: %v", err)
	}
	return np, nil
}

func (p PoliterType) Completed(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	if err := conj.Update(&conjunctor.ConjunctorUpdate{
		ReplacePrev: true,
		Inserts:     []conjunctor.Conjunctable{CompletedPoliter},
	}); err != nil {
		return nil, fmt.Errorf("Politer.Completed: %v", err)
	}
	return CompletedPoliter, nil
}

func (p PoliterType) Volitionally(conj *conjunctor.Conjunctor) (extender.Extender, error) {
	if err := conj.Update(&conjunctor.ConjunctorUpdate{
		ReplacePrev: true,
		Inserts:     []conjunctor.Conjunctable{VolitionalPoliter},
	}); err != nil {
		return nil, fmt.Errorf("Politer.Volitionally: %v", err)
	}
	return VolitionalPoliter, nil
}

type NegativePoliter struct {
	extender.UnimplementedExtender
}

func (p NegativePoliter) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Verb && conj.GetConjugationKind() != conjugationkind.Conjunctive {
		return nil, fmt.Errorf("Politer.OnAppend: cannot conjunct Politer to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}
	return &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AuxVerb,
		ConjugationKind: conjugationkind.Unknown,
		Inserts:         []conjunctor.Conjunctable{p},
	}, nil
}

func (p NegativePoliter) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, word.MustWord("ません", "ません")), nil
}

type CompletedPoliterType struct {
	extender.UnimplementedExtender
}

var CompletedPoliter = CompletedPoliterType{
	UnimplementedExtender: extender.NewUnimplementedExtender("completed politer aux verb"),
}

func (c CompletedPoliterType) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, word.MustWord("ました", "")), nil
}

func (c CompletedPoliterType) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Verb && conj.GetConjugationKind() != conjugationkind.Conjunctive {
		return nil, fmt.Errorf("CompletedPoliter.OnConjunct: cannot conjunct Politer to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}

	return &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AuxVerb,
		ConjugationKind: conjugationkind.Unknown,
		Inserts:         []conjunctor.Conjunctable{c},
	}, nil
}

type VolitionalPoliterType struct {
	extender.UnimplementedExtender
}

var VolitionalPoliter = VolitionalPoliterType{
	UnimplementedExtender: extender.NewUnimplementedExtender("volitional politer"),
}

func (v VolitionalPoliterType) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, word.MustWord("ましょう", "")), nil
}

func (v VolitionalPoliterType) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Verb && conj.GetConjugationKind() != conjugationkind.Conjunctive {
		return nil, fmt.Errorf("CompletedPoliter.OnConjunct: cannot conjunct Politer to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}

	return &conjunctor.ConjunctorUpdate{
		WordKind:        wordkind.AuxVerb,
		ConjugationKind: conjugationkind.Volitional,
		Inserts:         []conjunctor.Conjunctable{v},
	}, nil
}
