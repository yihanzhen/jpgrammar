package particle

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Particle struct {
	word.Word
	Extender extender.UnimplementedExtender
}

var Topic = Particle{
	Word: word.MustWord("は", ""),
}

var State = Particle{
	Word: word.MustWord("で", ""),
}

var Uncertainty = Particle{
	Word: word.MustWord("か", ""),
}

func (p Particle) OnAppend(conj *conjunctor.Conjunctor) error {
	if (conj.GetWordKind() != wordkind.Noun || conj.GetWordKind() != wordkind.Unknown) && conj.GetConjugationKind() != conjugationkind.Unknown {
		return fmt.Errorf("Particle.OnAppend: cannot conjunct particle to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}
	conj.UpdateWordKind(wordkind.Particle)
	conj.Insert(p)
	return nil
}

func (p Particle) OnWrite(_ conjunctor.Conjunctable, words []word.Word) ([]word.Word, error) {
	return append(words, p.Word), nil
}
