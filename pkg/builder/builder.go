package builder

import (
	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/builder/vocabulary"
)

type Builder struct {
	Vocab      vocabulary.Vocabulary
	Conjunctor conjunctor.Conjunctor
	Extender   extender.Extender
}

func (b *Builder) Append(text string) *Builder {
	w, err := b.Vocab.GetWord(text)
	if err != nil {
		return b
	}
	b.Conjunctor.Conjunct(w)
	return b
}
