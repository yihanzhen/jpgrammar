package verb

import "github.com/yihanzhen/jpgrammar/pkg/word"

type VerbConjugator interface {
	Imperfective() (word.Word, error)
	Adverbial() (word.Word, error)
	Attributive() (word.Word, error)
	Terminated() (word.Word, error)
	Hyperthetical() (word.Word, error)
	Volitional() (word.Word, error)
	Imperative() (word.Word, error)
}

type GodanVerbConjugator struct {
	verb Verb
}

func (c *GodanVerbConjugator) Imperfective() (word.Word, error) {
	w, err := word.LastKanaToCol(c.verb.Word, 0)
	if err != nil {
		return word.Word{}, err
	}

	return word.LastKanaAToWa(w), nil
}

func (c *GodanVerbConjugator) Adverbial() (word.Word, error) {
	return word.LastKanaToCol(c.verb.Word, 1)
}

func (c *GodanVerbConjugator) Attributive() (word.Word, error) {
	return word.LastKanaToCol(c.verb.Word, 2)
}

func (c *GodanVerbConjugator) Terminated() (word.Word, error) {
	return word.LastKanaToCol(c.verb.Word, 2)
}

func (c *GodanVerbConjugator) Hyperthetical() (word.Word, error) {
	return word.LastKanaToCol(c.verb.Word, 3)
}

func (c *GodanVerbConjugator) Imperative() (word.Word, error) {
	return word.LastKanaToCol(c.verb.Word, 4)
}

func (c *GodanVerbConjugator) Volitional() (word.Word, error) {
	return word.LastKanaToCol(c.verb.Word, 4)
}

type IchidanConjugator struct {
	verb Verb
}

func (c *IchidanConjugator) Imperfective() (word.Word, error) {
	return word.TrimLastKana(c.verb.Word)
}

func (c *IchidanConjugator) Adverbial() (word.Word, error) {
	return word.TrimLastKana(c.verb.Word)
}

func (c *IchidanConjugator) Attributive() (word.Word, error) {
	return word.LastKanaToCol(c.verb.Word, 2)
}

func (c *IchidanConjugator) Terminated() (word.Word, error) {
	return word.LastKanaToCol(c.verb.Word, 2)
}

func (c *IchidanConjugator) Hyperthetical() (word.Word, error) {
	return word.LastKanaToCol(c.verb.Word, 3)
}

func (c *IchidanConjugator) Imperative() (word.Word, error) {
	w, err := word.TrimLastKana(c.verb.Word)
	if err != nil {
		return word.Word{}, err
	}
	word.AppendSuffix(w, "ろ")
	return w, nil
}

func (c *IchidanConjugator) Volitional() (word.Word, error) {
	w, err := word.TrimLastKana(c.verb.Word)
	if err != nil {
		return word.Word{}, err
	}
	word.AppendSuffix(w, "よう")
	return w, nil
}
