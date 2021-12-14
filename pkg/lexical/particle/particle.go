package particle

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Particle struct {
	word.Word
	Extender extender.UnimplementedExtender
}

var Topic = Particle{
	Word:     word.MustWord("は", ""),
	Extender: extender.NewUnimplementedExtender("topic particle"),
}

var State = Particle{
	Word:     word.MustWord("で", ""),
	Extender: extender.NewUnimplementedExtender("state particle"),
}

var Uncertainty = Particle{
	Word:     word.MustWord("か", ""),
	Extender: extender.NewUnimplementedExtender("uncertainty particle"),
}

var LikewiseTopic = Particle{
	Word:     word.MustWord("も", ""),
	Extender: extender.NewUnimplementedExtender("likewise topic particle"),
}

var Reason = Particle{
	Word:     word.MustWord("から", ""),
	Extender: extender.NewUnimplementedExtender("reason particle"),
}

func (p Particle) OnConjunct(conj *conjunctor.Conjunctor) (*conjunctor.ConjunctorUpdate, error) {
	if (conj.GetWordKind() != wordkind.Noun || conj.GetWordKind() != wordkind.Unknown) && conj.GetConjugationKind() != conjugationkind.Unknown {
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
