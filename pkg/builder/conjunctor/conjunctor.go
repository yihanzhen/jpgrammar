package conjunctor

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Conjunctable interface {
	OnAppend(conj *Conjunctor) error
	OnWrite(Conjunctable, []word.Word) ([]word.Word, error)
}

type Conjunctor struct {
	parts           []Conjunctable
	wordKind        wordkind.WordKind
	conjugationKind conjugationkind.ConjugationKind
}

func NewConjunctor() *Conjunctor {
	return &Conjunctor{}
}

func (c *Conjunctor) Write() (string, error) {
	var words []word.Word
	var res string
	for i, part := range c.parts {
		var p Conjunctable
		if i != 0 {
			p = c.parts[i-1]
		}
		w, err := part.OnWrite(p, words)
		if err != nil {
			return "", fmt.Errorf("Conjunctor.Write: %v", err)
		}
		words = w
	}
	for _, w := range words {
		res = res + w.Write()
	}
	return res, nil
}

func (c *Conjunctor) RemoveHead() error {
	if len(c.parts) <= 1 {
		return fmt.Errorf("RemoveHead: nothing to remove")
	}
	c.parts = c.parts[0 : len(c.parts)-1]
	return nil
}

func (c *Conjunctor) Insert(part Conjunctable) {
	c.parts = append(c.parts, part)
}

func (c *Conjunctor) GetWordKind() wordkind.WordKind {
	return c.wordKind
}

func (c *Conjunctor) UpdateWordKind(wk wordkind.WordKind) {
	c.wordKind = wk
}

func (c *Conjunctor) UpdateConjugationKind(ck conjugationkind.ConjugationKind) {
	c.conjugationKind = ck
}

func (c *Conjunctor) GetConjugationKind() conjugationkind.ConjugationKind {
	return c.conjugationKind
}

func (c *Conjunctor) Append(parts ...Conjunctable) error {
	for _, p := range parts {
		if err := p.OnAppend(c); err != nil {
			return fmt.Errorf("Conjunctor.Append: %v", err)
		}
	}
	return nil
}
