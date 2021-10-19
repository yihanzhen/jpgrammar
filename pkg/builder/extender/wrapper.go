package extender

import (
	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/diag"
)

type ExtenderWrapper struct {
	Diag       *diag.Diag
	Conjunctor *conjunctor.Conjunctor

	Extender Extender
}

func NewExtenderWrapper(d *diag.Diag, c *conjunctor.Conjunctor) *ExtenderWrapper {
	return &ExtenderWrapper{
		Diag:       d,
		Conjunctor: c,
	}
}

func (w *ExtenderWrapper) Negated() *ExtenderWrapper {
	return w.wrapExtenderFunc(func() (Extender, error) {
		return w.Extender.Negated(w.Conjunctor)
	})
}

func (w *ExtenderWrapper) wrapExtenderFunc(f func() (Extender, error)) *ExtenderWrapper {
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
