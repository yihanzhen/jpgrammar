package particle

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/casekind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

// CaseMarker represents a case marker.
type CaseMarker struct {
	Case casekind.CaseKind
	Word word.Word
	extender.UnimplementedExtender
}

// OnConjunct implements the Conjunctable interface.
func (cm CaseMarker) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if conj.GetWordKind() != wordkind.Noun && conj.GetConjugationKind() != conjugationkind.Unknown {
		return nil, fmt.Errorf("CaseMarker.OnConjunct: cannot conjunct CaseMarker with word kind %v and conjugation kind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}
	if conj.GetCaseKind() != casekind.Unknown {
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

func newCaseMarker(writing, name string, caseKind casekind.CaseKind) CaseMarker {
	return CaseMarker{
		Case:                  caseKind,
		Word:                  word.MustWord(writing, ""),
		UnimplementedExtender: extender.NewUnimplementedExtender(name),
	}
}

// From returns a CaseMarker given a case kind.
func From(caseKind casekind.CaseKind) (CaseMarker, error) {
	marker, exists := caseMarkerMap[caseKind]
	if !exists {
		return CaseMarker{}, fmt.Errorf("From: unknown case kind: %v", caseKind)
	}
	return marker, nil
}

var caseMarkerMap map[casekind.CaseKind]CaseMarker = map[casekind.CaseKind]CaseMarker{
	casekind.Start:     newCaseMarker("から", "start case marker", casekind.Start),
	casekind.End:       newCaseMarker("まで", "end case marker", casekind.End),
	casekind.Timestamp: newCaseMarker("に", "timestamp case marker", casekind.Timestamp),
	casekind.Time:      {Case: casekind.Time, Word: word.Omitted, UnimplementedExtender: extender.NewUnimplementedExtender("time case marker")},
}
