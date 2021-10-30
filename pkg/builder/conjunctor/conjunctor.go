package conjunctor

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation/kind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

// Conjunctable represents any component that can be conjuncted in a sentence.
type Conjunctable interface {
	OnConjunct(conj *Conjunctor) (*ConjunctorUpdate, error)
	OnWrite([]word.Word, ...Conjunctable) ([]word.Word, error)
}

// conjunctableWrapper wraps a Conjunctable with additional states during conjunction.
type conjunctableWrapper struct {
	conjunctable Conjunctable
	cached       bool
}

// Conjunctor conjuncts components to form a sentance.
type Conjunctor struct {
	parts           []conjunctableWrapper
	wordKind        wordkind.WordKind
	conjugationKind kind.ConjugationKind
}

// ConjunctorUpdate is the callback response of OnConjunct to update the state
// of the Conjunctor.
type ConjunctorUpdate struct {
	WordKind        wordkind.WordKind
	ConjugationKind kind.ConjugationKind
	CachePrev       bool
	ReplacePrev     bool
	Inserts         []Conjunctable
}

// NewConjunctor creates a new conjunctor.
func NewConjunctor() *Conjunctor {
	return &Conjunctor{}
}

// Write returns the sentence constructed.
func (c *Conjunctor) Write() (string, error) {
	var words []word.Word
	var cache []Conjunctable
	for _, part := range c.parts {
		if part.cached {
			cache = append(cache, part.conjunctable)
			continue
		}
		w, err := part.conjunctable.OnWrite(words, cache...)
		if err != nil {
			return "", fmt.Errorf("Conjunctor.Write: %v", err)
		}
		cache = []Conjunctable{}
		words = w
	}

	var res string
	for _, w := range words {
		res = res + w.Write()
	}
	return res, nil
}

// GetWordKind returns the latest word kind.
func (c *Conjunctor) GetWordKind() wordkind.WordKind {
	return c.wordKind
}

// GetWordKind returns the latest conjugation kind.
func (c *Conjunctor) GetConjugationKind() kind.ConjugationKind {
	return c.conjugationKind
}

// Conjunct conjuncts one or more Conjunctables.
func (c *Conjunctor) Conjunct(parts ...Conjunctable) error {
	for _, p := range parts {
		cu, err := p.OnConjunct(c)
		if err != nil {
			return fmt.Errorf("Conjunctor.Conjunct: %v", err)
		}
		if err := c.Update(cu); err != nil {
			return fmt.Errorf("Conjunctor.Conjunct: %v", err)
		}
	}
	return nil
}

func (c *Conjunctor) Update(cu *ConjunctorUpdate) error {
	c.wordKind = cu.WordKind
	c.conjugationKind = cu.ConjugationKind
	if cu.CachePrev {
		if len(c.parts) == 0 {
			return fmt.Errorf("got ConjunctorUpdate CachePrev == true, but Conjunctor has no prev")
		}
		c.parts[len(c.parts)-1].cached = true
	}
	if cu.ReplacePrev {
		if len(c.parts) == 0 {
			return fmt.Errorf("got ConjunctorUpdate ReplacePrev == true, but Conjunctor has no prev")
		}
		c.parts = c.parts[:len(c.parts)-1]
	}
	for _, insert := range cu.Inserts {
		c.parts = append(c.parts, conjunctableWrapper{conjunctable: insert})
	}
	return nil
}
