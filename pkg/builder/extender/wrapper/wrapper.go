package wrapper

import (
	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/diag"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/builder/vocabulary"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/casekind"
)

type ExtenderWrapper struct {
	Vocab      *vocabulary.Vocabulary
	Diag       *diag.Diag
	Conjunctor *conjunctor.Conjunctor

	Extender extender.Extender
}

func NewExtenderWrapper(v *vocabulary.Vocabulary, d *diag.Diag, c *conjunctor.Conjunctor) *ExtenderWrapper {
	return &ExtenderWrapper{
		Vocab:      v,
		Diag:       d,
		Conjunctor: c,
	}
}

func (w *ExtenderWrapper) Negated() *ExtenderWrapper {
	return w.wrapExtenderFunc(func() (extender.Extender, error) {
		return w.Extender.Negated(w.Conjunctor)
	})
}

func (w *ExtenderWrapper) Asserted() *ExtenderWrapper {
	return w.wrapExtenderFunc(func() (extender.Extender, error) {
		return w.Extender.Asserted(w.Conjunctor)
	})
}

func (w *ExtenderWrapper) Attributing(noun string) *ExtenderWrapper {
	n, err := w.Vocab.GetWord(noun)
	if err != nil {
		w.Diag.SaveError(err)
		return w
	}
	ext, err := w.Extender.Attributing(w.Conjunctor, n)
	if err != nil {
		w.Diag.SaveError(err)
		return w
	}
	w.Extender = ext
	return w
}

func (w *ExtenderWrapper) As(ck casekind.CaseKind) *ExtenderWrapper {
	ext, err := w.Extender.As(w.Conjunctor, ck)
	if err != nil {

	}
	w.Extender = ext
	return w
}

func (w *ExtenderWrapper) Politely() *ExtenderWrapper {
	ext, err := w.Extender.Politely(w.Conjunctor)
	if err != nil {

	}
	w.Extender = ext
	return w
}

func (w *ExtenderWrapper) Completed() *ExtenderWrapper {
	ext, err := w.Extender.Completed(w.Conjunctor)
	if err != nil {

	}
	w.Extender = ext
	return w
}

func (w *ExtenderWrapper) wrapExtenderFunc(f func() (extender.Extender, error)) *ExtenderWrapper {
	if w.Diag.HasErrors() {
		return w
	}
	e, err := f()
	if err != nil {
		w.Diag.SaveError(err)
		return w
	}
	w.Extender = e
	return w
}
