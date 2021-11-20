package conjugation

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	ck "github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	wk "github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Conjugatable interface {
	Conjugate(ck.ConjugationKind) (word.Word, error)
}

// A conjugation represents a conjugation for a verb, adjective, adjectival noun, noun or
// aux verb.
type Conjugation struct {
	kind ck.ConjugationKind
}

// NewConjugation returns a new conjugation.
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
		ck.Irrealis:    true,
		ck.Conjunctive: true,
		ck.Attributive: true,
		ck.Terminal:    true,
		ck.Conditional: true,
		ck.Volitional:  true,
	},
	wk.Adjective: {
		ck.Irrealis:    true,
		ck.Conjunctive: true,
		ck.Attributive: true,
		ck.Terminal:    true,
		ck.Conditional: true,
	},
	wk.AdjNoun: {
		ck.Attributive: true,
	},
}

// OnConjunct implements the Conjunctable interface.
func (cj Conjugation) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	cjs, ok := conjugationPrevTable[conj.GetWordKind()]
	if !ok {
		return nil, fmt.Errorf("Conjugation.OnAppend: unknown wordkind for conjugation %v", cj)
	}
	if _, ok := cjs[cj.kind]; !ok {
		return nil, fmt.Errorf("Conjugation.OnAppend: %v does not have conjugation %v", conj.GetWordKind(), cj)
	}

	return &conjunctor.ConjunctorUpdate{
		ConjugationKind: cj.kind,
		Inserts:         []conjunctor.Conjunctable{cj},
		CachePrev:       true,
	}, nil
}

// OnWrite implements the Conjunctable interface.
func (cj Conjugation) OnWrite(words []word.Word, prev ...conjunctor.Conjunctable) ([]word.Word, error) {
	if len(prev) != 1 {
		return nil, fmt.Errorf("Conjugation(%s).OnWrite: got prev %v, want only one", cj.kind, prev)
	}
	if p, ok := prev[0].(Conjugatable); ok {
		w, err := p.Conjugate(cj.kind)
		if err != nil {
			return nil, fmt.Errorf("Conjugation(%s).OnWrite: %v", cj.kind, err)
		}
		words = append(words, w)
		return words, nil
	}
	return nil, fmt.Errorf("Conjugation.OnWrite(%s): Conjunctable not conjugatble", cj.kind)
}
