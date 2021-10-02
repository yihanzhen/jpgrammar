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

}
