package builder

import (
	"fmt"

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
	e := extender.UnimplementedExtender{
		SaveError: func(err error) {
			d.SaveError(err)
		},
		GetError: func() error {
			if !d.HasErrors() {
				return nil
			}
			return fmt.Errorf("errors: %v", d.GetErrors())
		},
	}
	b := Builder{
		Vocab:      vocabulary.NewVocabulary(),
		Conjunctor: conjunctor.NewConjugator(),
		Extender:   e,
		Diag:       d,
	}
	return &b
}

func (b *Builder) Append(text string) *Builder {
	w, err := b.Vocab.GetWord(text)
	if err != nil {
		b.Diag.SaveError(err)
		return b
	}
	b.Conjunctor.Conjunct(w)
	return b
}
