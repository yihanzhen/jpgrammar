package conjugation

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	ck "github.com/yihanzhen/jpgrammar/pkg/lexical/conjugationkind"
	wk "github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Conjugatable interface {
	Conjugate(ck.ConjugationKind) (word.Word, error)
}

type Conjugation struct {
	kind ck.ConjugationKind
}

func NewConjugation(kind ck.ConjugationKind) Conjugation {
	return Conjugation{
		kind: kind,
	}
}

var conjugationPrevTable map[wk.WordKind]map[ck.ConjugationKind]bool = map[wk.WordKind]map[ck.ConjugationKind]bool{
	wk.Noun: {
		ck.Attributive: true,
	},
	wk.Verb: {
		ck.Imperfective: true,
		ck.Conjunctive:  true,
		ck.Attributive:  true,
		ck.Terminal:     true,
		ck.Conditional:  true,
		ck.Volitional:   true,
	},
	wk.Adjective: {
		ck.Imperfective: true,
		ck.Conjunctive:  true,
		ck.Attributive:  true,
		ck.Terminal:     true,
		ck.Conditional:  true,
	},
	wk.AdjNoun: {
		ck.Attributive: true,
	},
}

func (cj Conjugation) OnAppend(conj *conjunctor.Conjunctor) error {
	cjs, ok := conjugationPrevTable[conj.GetWordKind()]
	if !ok {
		return fmt.Errorf("Conjugation.OnAppend: unknown wordkind for conjugation %v", cj)
	}
	if _, ok := cjs[cj.kind]; !ok {
		return fmt.Errorf("Conjugation.OnAppend: %v does not have conjugation %v", conj.GetWordKind(), cj)
	}
	conj.Insert(cj)
	conj.UpdateConjugationKind(cj.kind)
	return nil
}

func (cj Conjugation) OnWrite(prev conjunctor.Conjunctable, words []word.Word) ([]word.Word, error) {
	if p, ok := prev.(Conjugatable); ok {
		words = words[:len(words)-1]
		w, err := p.Conjugate(cj.kind)
		if err != nil {
			return nil, fmt.Errorf("Conjugation.OnWrite: %v", err)
		}
		words = append(words, w)
		return words, nil
	}
	return nil, fmt.Errorf("Conjugation.OnWrite: Conjunctable not conjugatble")
}
