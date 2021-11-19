package particle

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/casing"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation/kind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type CaseMarker struct {
	Case casing.CaseKind
	Word word.Word
	extender.UnimplementedExtender
}

func (cm CaseMarker) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Noun && conj.GetConjugationKind() != kind.Unknown {
		return nil, fmt.Errorf("CaseMarker.OnConjunct: cannot conjunct CaseMarker with word kind %v and conjugation kind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}
	if conj.GetCaseKind() != casing.Unknown {
		return nil, fmt.Errorf("CaseMarker.OnConjunct: already has case marker: %v", conj.GetCaseKind())
	}
	return &conjunctor.ConjunctorUpdate{
		Case:     cm.Case,
		Inserts:  []conjunctor.Conjunctable{cm},
		WordKind: conj.GetWordKind(),
	}, nil
}

// OnWrite implements the Conjunctable interface.
func (cm CaseMarker) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, cm.Word), nil
}

func newCaseMarker(writing, name string, caseKind casing.CaseKind) CaseMarker {
	return CaseMarker{
		Case:                  caseKind,
		Word:                  word.MustWord(writing, ""),
		UnimplementedExtender: extender.NewUnimplementedExtender(name),
	}
}

func From(caseKind casing.CaseKind) (CaseMarker, error) {
	marker, exists := caseMarkerMap[caseKind]
	if !exists {
		return CaseMarker{}, fmt.Errorf("From: unknown case kind: %v", caseKind)
	}
	return marker, nil
}

var caseMarkerMap map[casing.CaseKind]CaseMarker = map[casing.CaseKind]CaseMarker{
	casing.Start:     newCaseMarker("から", "start case marker", casing.Start),
	casing.End:       newCaseMarker("まで", "end case marker", casing.Start),
	casing.Timestamp: newCaseMarker("に", "timestamp case marker", casing.Start),
	casing.Time:      {Case: casing.Time, Word: word.Omitted, UnimplementedExtender: extender.NewUnimplementedExtender("time case marker")},
}
