package verb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type verbConjugator interface {
	Imperfective() (word.Word, error)
	Conjunctive() (word.Word, error)
	Attributive() (word.Word, error)
	Terminal() (word.Word, error)
	Conditional() (word.Word, error)
	Volitional() (word.Word, error)
	Imperative() (word.Word, error)
}

func (v Verb) getConjugator() verbConjugator {
	if v.forceTypeOneConj {
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

func (c *TypeOneVerbConjugator) Imperfective() (word.Word, error) {
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

func (c *TypeOneVerbConjugator) Conjunctive() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(1))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Adverbial: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) Attributive() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(2))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Attributive: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) Terminal() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(2))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Terminated: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) Conditional() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(3))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Hyperthetical: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) Imperative() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(3))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Imperative: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) Volitional() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(4))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Imperative: %v", err)
	}
	return w, nil
}

type TypeTwoVerbConjugator struct {
	verb Verb
}

func (c *TypeTwoVerbConjugator) Imperfective() (word.Word, error) {
	w, err := c.verb.Word.TrimLastRune()
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeTwoVerbConjugator.Imperfective: %v", err)
	}
	return w, nil
}

func (c *TypeTwoVerbConjugator) Conjunctive() (word.Word, error) {
	w, err := c.verb.Word.TrimLastRune()
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeTwoVerbConjugator.Adverbial: %v", err)
	}
	return w, nil
}

func (c *TypeTwoVerbConjugator) Attributive() (word.Word, error) {
	return c.verb.Word, nil
}

func (c *TypeTwoVerbConjugator) Terminal() (word.Word, error) {
	return c.verb.Word, nil
}

func (c *TypeTwoVerbConjugator) Conditional() (word.Word, error) {
	return c.verb.Word, nil
}

func (c *TypeTwoVerbConjugator) Imperative() (word.Word, error) {
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

func (c *TypeTwoVerbConjugator) Volitional() (word.Word, error) {
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
