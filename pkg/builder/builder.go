package builder

import (
	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/diag"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/builder/vocabulary"
)

type Builder struct {
	Vocab      *vocabulary.Vocabulary
	Conjunctor *conjunctor.Conjunctor
	Extender   *extender.ExtenderWrapper
	Diag       *diag.Diag
}

func NewBuilder() *Builder {
	d := &diag.Diag{}
	c := conjunctor.NewConjunctor()
	e := extender.NewExtenderWrapper(d, c)
	b := Builder{
		Vocab:      vocabulary.NewVocabulary(),
		Conjunctor: c,
		Extender:   e,
		Diag:       d,
	}
	return &b
}

func (b *Builder) Append(text string) *Builder {
	if b.Diag.HasErrors() {
		return b
	}
	w, err := b.Vocab.GetWord(text)
	if err != nil {
		b.Diag.SaveError(err)
		return b
	}
	if err := b.Conjunctor.Append(w); err != nil {
		b.Diag.SaveError(err)
	}
	return b
}

func (b *Builder) Build() (string, error) {
	return b.Conjunctor.Write()
}
