package verb

import (
	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
)

type VerbExtender struct {
	extender.UnimplementedExtender
	conj *conjunctor.Conjunctor
}

func (v *VerbExtender) Politely() extender.Extender {
	return v
}
