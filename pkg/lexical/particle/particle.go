package particle

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation/kind"
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

var LikewiseTopic = Particle{
	Word: word.MustWord("も", ""),
}

func (p Particle) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if (conj.GetWordKind() != wordkind.Noun || conj.GetWordKind() != wordkind.Unknown) && conj.GetConjugationKind() != kind.Unknown {
		return nil, fmt.Errorf("Particle.OnAppend: cannot conjunct particle to wordkind %v and conjugationkind %v", conj.GetWordKind(), conj.GetConjugationKind())
	}

	return &conjunctor.ConjunctorUpdate{
		WordKind: wordkind.Particle,
		Inserts:  []conjunctor.Conjunctable{p},
	}, nil
}

func (p Particle) OnWrite(words []word.Word, _ ...conjunctor.Conjunctable) ([]word.Word, error) {
	return append(words, p.Word), nil
}
