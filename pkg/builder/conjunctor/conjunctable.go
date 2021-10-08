package conjunctor

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Conjunctable interface {
	CheckPrev(*Conjunctor, Conjunctable) error
	CheckNext(*Conjunctor, Conjunctable) error
	OnConjunct(prev, next Conjunctable) ([]Conjunctable, error)
	OnWrite([]word.Word) []word.Word

	GetWordKind() wordkind.WordKind
	GetConjugationKind() conjugationkind.ConjugationKind
}

type DefaultConjunctable struct{}

func (d DefaultConjunctable) CheckPrev(*Conjunctor, Conjunctable) error {
	return nil
}

func (d DefaultConjunctable) CheckNext(*Conjunctor, Conjunctable) error {
	return nil
}

func (d DefaultConjunctable) OnConjunct(prev, next Conjunctable) ([]Conjunctable, error) {
	return []Conjunctable{prev, d, next}, nil
}

func (d DefaultConjunctable) OnWrite(words []word.Word) []word.Word {
	return words
}

func (d DefaultConjunctable) GetWordKind() wordkind.WordKind {
	return wordkind.Unknown
}

func (d DefaultConjunctable) GetConjugationKind() conjugationkind.ConjugationKind {
	return conjugationkind.Unknown
}

type SentenceStartType struct {
	DefaultConjunctable
}

func (s SentenceStartType) CheckPrev(c *Conjunctor, p Conjunctable) error {
	return fmt.Errorf("SentenceStart.CheckPrev: should never be called for SentenceStart")
}

func (s SentenceStartType) CheckNext(c *Conjunctor, p Conjunctable) error {
	return nil
}

func (s SentenceStartType) OnWrite(words []word.Word) []word.Word {
	return words
}

func (s SentenceStartType) OnConjunct(_, _ Conjunctable) ([]Conjunctable, error) {
	return nil, fmt.Errorf("SentenceStart.OnConjunct: should never be called for SentenceStart")
}

type SentenceEndType struct {
	DefaultConjunctable
}

func (s SentenceEndType) CheckPrev(c *Conjunctor, p Conjunctable) error {
	return nil
}

func (s SentenceEndType) CheckNext(c *Conjunctor, p Conjunctable) error {
	return fmt.Errorf("SentenceEnd.CheckNext: should never be called for SentenceEnd")
}

func (s SentenceEndType) OnWrite(words []word.Word) []word.Word {
	return words
}

func (s SentenceEndType) OnConjunct(_, _ Conjunctable) ([]Conjunctable, error) {
	return nil, fmt.Errorf("SentenceEnd.OnConjunct: should never be called for SentenceEnd")
}

var (
	SentenceStart = SentenceStartType{}
	SentenceEnd   = SentenceEndType{}
)
