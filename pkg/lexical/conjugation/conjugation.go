package conjugation

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	ck "github.com/yihanzhen/jpgrammar/pkg/lexical/conjugationkind"
	wk "github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
)

type Conjugation struct {
	kind ck.ConjugationKind
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

func (cj Conjugation) CheckPrev(c *conjunctor.Conjunctor, p conjunctor.Conjunctable) error {
	cjs, ok := conjugationPrevTable[p.GetWordKind()]
	if !ok {
		return fmt.Errorf("Conjugation.OnConjunctPrev: unknown wordkind for conjugation %v", cj)
	}
	if _, ok := cjs[cj.kind]; !ok {
		return fmt.Errorf("Conjugation.OnConjunctPrev: %v does not have conjugation %v", p.GetWordKind(), cj)
	}
	return nil
}

func (cj Conjugation) CheckNext(c *conjunctor.Conjunctor, p conjunctor.Conjunctable) error {
	return nil
}

func (cj Conjugation) OnConjunct()
