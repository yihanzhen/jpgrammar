package auxverb

import (
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type AuxVerb struct {
	word.Word
	Extender extender.Extender
}

var PoliteAssertor = AuxVerb{
	Extender: &PoliteAssertorExtender{},
}

type PoliteAssertorExtender struct {
	extender.UnimplementedExtender
}

func (e *PoliteAssertorExtender) Negated() extender.Extender {
	return nil
}
