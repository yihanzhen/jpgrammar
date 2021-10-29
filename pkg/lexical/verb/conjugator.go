package verb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

func (v Verb) Conjugate(ck conjugationkind.ConjugationKind) (word.Word, error) {
	conj := v.getConjugator()
	switch ck {
	case conjugationkind.Imperfective:
		return conj.imperfective()
	case conjugationkind.Conjunctive:
		return conj.conjunctive()
	case conjugationkind.Attributive:
		return conj.attributive()
	case conjugationkind.Terminal:
		return conj.terminal()
	case conjugationkind.Conditional:
		return conj.conditional()
	case conjugationkind.Volitional:
		return conj.volitional()
	}
	return word.Word{}, fmt.Errorf("Verb.Conjugate: conjugationKind not Conjugatable: %v", ck)
}

type verbConjugator interface {
	imperfective() (word.Word, error)
	conjunctive() (word.Word, error)
	attributive() (word.Word, error)
	terminal() (word.Word, error)
	conditional() (word.Word, error)
	volitional() (word.Word, error)
	imperative() (word.Word, error)
}

func (v Verb) getConjugator() verbConjugator {
	if v.forceConjTypeOne {
		return &TypeOneVerbConjugator{
			verb: v,
		}
	}
	if !v.CheckLastRune('る') {
		return &TypeOneVerbConjugator{
			verb: v,
		}
	}
	if !v.CheckRune(word.NthLastRune(1), word.IsCol(1)) && !v.CheckRune(word.NthLastRune(1), word.IsCol(3)) {
		return &TypeOneVerbConjugator{
			verb: v,
		}
	}
	return &TypeTwoVerbConjugator{
		verb: v,
	}
}

type TypeOneVerbConjugator struct {
	verb Verb
}

func (c *TypeOneVerbConjugator) imperfective() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(0))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Imperfective: %v", err)
	}
	w, err = w.ChangeLastRuneTo(word.AToWa)
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Imperfective: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) conjunctive() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(1))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Adverbial: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) attributive() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(2))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Attributive: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) terminal() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(2))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Terminated: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) conditional() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(3))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Hyperthetical: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) imperative() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(3))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Imperative: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) volitional() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(4))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Imperative: %v", err)
	}
	return w, nil
}

type TypeTwoVerbConjugator struct {
	verb Verb
}

func (c *TypeTwoVerbConjugator) imperfective() (word.Word, error) {
	w, err := c.verb.Word.TrimLastRune()
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeTwoVerbConjugator.Imperfective: %v", err)
	}
	return w, nil
}

func (c *TypeTwoVerbConjugator) conjunctive() (word.Word, error) {
	w, err := c.verb.Word.TrimLastRune()
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeTwoVerbConjugator.Adverbial: %v", err)
	}
	return w, nil
}

func (c *TypeTwoVerbConjugator) attributive() (word.Word, error) {
	return c.verb.Word, nil
}

func (c *TypeTwoVerbConjugator) terminal() (word.Word, error) {
	return c.verb.Word, nil
}

func (c *TypeTwoVerbConjugator) conditional() (word.Word, error) {
	return c.verb.Word, nil
}

func (c *TypeTwoVerbConjugator) imperative() (word.Word, error) {
	w, err := c.verb.Word.TrimLastRune()
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeTwoVerbConjugator.Imperative: %v", err)
	}
	w, err = w.Append("ろ")
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeTwoVerbConjugator.Imperative: %v", err)
	}
	return w, nil
}

func (c *TypeTwoVerbConjugator) volitional() (word.Word, error) {
	w, err := c.verb.Word.TrimLastRune()
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeTwoVerbConjugator.Volitional: %v", err)
	}
	w, err = w.Append("よう")
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeTwoVerbConjugator.Volitional: %v", err)
	}
	return w, nil
}
