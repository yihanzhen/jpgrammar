package conjunctor

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Conjunctor struct {
	parts []Conjunctable
	pos   int
	err   error
}

func NewConjugator() *Conjunctor {
	return &Conjunctor{
		parts: []Conjunctable{SentenceStart, SentenceEnd},
		pos:   1,
	}
}

func (c *Conjunctor) Write() string {
	var words []word.Word
	var res string
	for _, part := range c.parts {
		words = part.OnWrite(words)
	}
	for _, w := range words {
		res = res + w.Write()
	}
	return res
}

func (c *Conjunctor) GetHead() Conjunctable {
	return c.parts[c.pos]
}

func (c *Conjunctor) MoveHead(ind int) {
	c.pos = ind
}

func (c *Conjunctor) Conjunct(part Conjunctable) {
	if c.err != nil {
		return
	}
	if c.pos <= 0 || c.pos >= len(c.parts)-1 {
		c.err = fmt.Errorf("invalid pos, got %v, want [1, %v)", c.pos, len(c.parts)-1)
		return
	}
	prev := c.parts[c.pos-1]
	next := c.parts[c.pos]
	conjuncted, err := part.OnConjunct(prev, next)
	if err != nil {
		c.err = fmt.Errorf("Conjunctor.Conjunct: %v", err)
	}
	var parts []Conjunctable
	parts = append(parts, c.parts[:c.pos-1]...)
	parts = append(parts, conjuncted...)
	parts = append(parts, c.parts[c.pos:]...)
	c.parts = parts
}

func (c *Conjunctor) GetError() error {
	return c.err
}
