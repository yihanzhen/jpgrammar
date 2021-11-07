package word

import (
	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
)

type Word interface {
	conjunctor.Conjunctable
	extender.Extender
}

const OmittedNoun = "omitted-noun"
