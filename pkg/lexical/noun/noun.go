package noun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Noun struct {
	word.Word
	extender.UnimplementedExtender
}

func (n Noun) OnAppend(conj *conjunctor.Conjunctor) error {
	// Starting a new sentence.
	if conj.GetWordKind() == wordkind.Unknown && conj.GetConjugationKind() == conjugationkind.Unknown {
		conj.UpdateWordKind(wordkind.Noun)
		conj.UpdateWordKind(wordkind.Unknown)
		conj.Insert(n)
		return nil
	}
	// Starting a new component.
	if conj.GetWordKind() == wordkind.Particle {
		conj.UpdateWordKind(wordkind.Noun)
		conj.UpdateWordKind(wordkind.Unknown)
		conj.Insert(n)
		return nil
	}

	// Compound words.
	if conj.GetWordKind() == wordkind.Noun && conj.GetConjugationKind() == conjugationkind.Unknown {
		conj.UpdateWordKind(wordkind.Noun)
		conj.UpdateWordKind(wordkind.Unknown)
		conj.Insert(n)
		return nil
	}

	if conj.GetWordKind() == wordkind.Verb && conj.GetConjugationKind() == conjugationkind.Conjunctive {
		conj.UpdateWordKind(wordkind.Noun)
		conj.UpdateWordKind(wordkind.Unknown)
		conj.Insert(n)
		return nil
	}

	// Attributive.
	if conj.GetConjugationKind() == conjugationkind.Attributive {
		conj.UpdateWordKind(wordkind.Noun)
		conj.UpdateWordKind(wordkind.Unknown)
		conj.Insert(n)
		return nil
	}

	return fmt.Errorf("Noun.OnAppend: cannot conjunct noun to wordkind %v and conjugationkind: %v", conj.GetWordKind(), conj.GetConjugationKind())
}

func (n Noun) OnWrite(_ conjunctor.Conjunctable, words []word.Word) ([]word.Word, error) {
	return append(words, n.Word), nil
}
