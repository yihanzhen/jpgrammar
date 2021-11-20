package verb

import (
	"fmt"

	kind "github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

func (v Verb) Conjugate(ck kind.ConjugationKind) (word.Word, error) {
	conj := v.getConjugator()
	switch ck {
	case kind.Irrealis:
		return conj.irrealis()
	case kind.Conjunctive:
		return conj.conjunctive()
	case kind.Attributive:
		return conj.attributive()
	case kind.Terminal:
		return conj.terminal()
	case kind.Conditional:
		return conj.conditional()
	case kind.Volitional:
		return conj.volitional()
	}
	return word.Word{}, fmt.Errorf("Verb.Conjugate: conjugationKind not Conjugatable: %v", ck)
}

type verbConjugator interface {
	irrealis() (word.Word, error)
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
	if v.CheckSuffix("する") {
		return &suruConjugator{
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

func (c *TypeOneVerbConjugator) irrealis() (word.Word, error) {
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

func (c *TypeTwoVerbConjugator) irrealis() (word.Word, error) {
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

type suruConjugator struct {
	verb Verb
}

func (c *suruConjugator) irrealis() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRunes(2, "し")
	if err != nil {
		return word.Word{}, fmt.Errorf("suruVerbConjugator.Imperfective: %v", err)
	}
	w, err = w.ChangeLastRuneTo(word.AToWa)
	if err != nil {
		return word.Word{}, fmt.Errorf("suruVerbConjugator.Imperfective: %v", err)
	}
	return w, nil
}

func (c *suruConjugator) conjunctive() (word.Word, error) {
	w, err := c.verb.Word.ChangeLastRunes(2, "し")
	if err != nil {
		return word.Word{}, fmt.Errorf("suruVerbConjugator.conjunctive: %v", err)
	}
	w, err = w.ChangeLastRuneTo(word.AToWa)
	if err != nil {
		return word.Word{}, fmt.Errorf("suruVerbConjugator.conjunctive: %v", err)
	}
	return w, nil
}

func (c *suruConjugator) attributive() (word.Word, error) {
	return word.Word{}, fmt.Errorf("unimplemented: suruVerbConjugator.attributive")
}

func (c *suruConjugator) terminal() (word.Word, error) {
	return word.Word{}, fmt.Errorf("unimplemented: suruVerbConjugator.terminal")

}

func (c *suruConjugator) conditional() (word.Word, error) {
	return word.Word{}, fmt.Errorf("unimplemented: suruVerbConjugator.conditional")

}

func (c *suruConjugator) imperative() (word.Word, error) {
	return word.Word{}, fmt.Errorf("unimplemented: suruVerbConjugator.imperitive")

}

func (c *suruConjugator) volitional() (word.Word, error) {
	return word.Word{}, fmt.Errorf("unimplemented: suruVerbConjugator.volitional")

}
