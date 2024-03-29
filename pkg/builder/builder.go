package builder

import (
	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/diag"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender/wrapper"
	"github.com/yihanzhen/jpgrammar/pkg/builder/vocabulary"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
)

// Builder represents a sentence builder.
type Builder struct {
	Vocab      *vocabulary.Vocabulary
	Conjunctor *conjunctor.Conjunctor
	Diag       *diag.Diag
	*wrapper.ExtenderWrapper
}

// NewBuilder creates a new Builder.
func NewBuilder() *Builder {
	v := vocabulary.NewVocabulary()
	d := &diag.Diag{}
	c := conjunctor.NewConjunctor()
	e := wrapper.NewExtenderWrapper(v, d, c)
	b := Builder{
		Vocab:           v,
		Conjunctor:      c,
		ExtenderWrapper: e,
		Diag:            d,
	}
	return &b
}

// Append looks up a word from vocabulary, and then append it to the Conjunctor.
// To chain calls to conjugate or extend the word appended, Append saves any error
// to Diag instead of returning them. All following calls are no-ops if Diag
// has any existing errors.
func (b *Builder) Append(text string) *Builder {
	if b.Diag.HasErrors() {
		return b
	}
	w, err := b.Vocab.GetWord(text)
	if err != nil {
		b.Diag.SaveError(err)
		return b
	}
	if err := b.Conjunctor.Conjunct(w); err != nil {
		b.Diag.SaveError(err)
	}
	b.ExtenderWrapper.Extender = w
	return b
}

// Mark appends a particle to the Conjunctor.
// To chain calls to conjugate or extend the word appended, Make saves any error
// to Diag instead of returning them. All following calls are no-ops if Diag
// has any existing errors.
func (b *Builder) Mark(p particle.Particle) *Builder {
	if b.Diag.HasErrors() {
		return b
	}
	if err := b.Conjunctor.Conjunct(p); err != nil {
		b.Diag.SaveError(err)
	}
	return b
}

func (b *Builder) WithContraction(ct interface{}) {

}

// Build returns the sentence built.
func (b *Builder) Build() (string, error) {
	return b.Conjunctor.Write()
}
