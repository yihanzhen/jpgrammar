package adjective

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

func CompletedEuphony(w word.Word) (word.Word, error) {
	if !w.CheckSuffix("く") {
		return word.Word{}, fmt.Errorf("adjective.CompleteEuphony: must end with く")
	}
	w2, err := w.ChangeLastRune('か', word.IgnoreConjugateRef)
	if err != nil {
		return word.Word{}, fmt.Errorf("adjective.CompleteEuphony: %v", err)
	}
	w3, err := w2.Append("っ")
	if err != nil {
		return word.Word{}, fmt.Errorf("adjective.CompleteEuphony: %v", err)
	}
	return w3, nil
}

func (adj Adjective) Conjugate(conj conjugationkind.ConjugationKind) (word.Word, error) {
	switch conj {
	case conjugationkind.Conjunctive:
		w2, err := adj.Word.ChangeLastRune('く', word.IgnoreConjugateRef)
		if err != nil {
			return w2, fmt.Errorf("adjective.Conjugate: %v", err)
		}
		return w2, nil
	case conjugationkind.Attributive:
		return adj.Word, nil
	}
	return word.Word{}, fmt.Errorf("unknown conjugation kind: %s", conj)
}
