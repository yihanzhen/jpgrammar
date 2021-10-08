package verb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type VerbConjugator interface {
	Imperfective() (word.Word, error)
	Adverbial() (word.Word, error)
	Attributive() (word.Word, error)
	Terminated() (word.Word, error)
	Hyperthetical() (word.Word, error)
	Volitional() (word.Word, error)
	Imperative() (word.Word, error)
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

func (c *TypeOneVerbConjugator) Adverbial() (word.Word, error) {
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

func (c *TypeOneVerbConjugator) Terminated() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRuneTo(word.ToCol(2))
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeOneVerbConjugator.Terminated: %v", err)
	}
	return w, nil
}

func (c *TypeOneVerbConjugator) Hyperthetical() (word.Word, error) {
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

func (c *TypeTwoVerbConjugator) Adverbial() (word.Word, error) {
	w, err := c.verb.Word.TrimLastRune()
	if err != nil {
		return word.Word{}, fmt.Errorf("TypeTwoVerbConjugator.Adverbial: %v", err)
	}
	return w, nil
}

func (c *TypeTwoVerbConjugator) Attributive() (word.Word, error) {
	return c.verb.Word, nil
}

func (c *TypeTwoVerbConjugator) Terminated() (word.Word, error) {
	return c.verb.Word, nil
}

func (c *TypeTwoVerbConjugator) Hyperthetical() (word.Word, error) {
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
