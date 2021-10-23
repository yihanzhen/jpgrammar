package extender

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
)

type Extender interface {
	// AsSubject(*conjunctor.Conjunctor) (Extender, error)
	// AsObject(*conjunctor.Conjunctor) (Extender, error)
	// AsDirection(*conjunctor.Conjunctor) (Extender, error)
	// AsVenue(*conjunctor.Conjunctor) (Extender, error)
	// AsTimestamp(*conjunctor.Conjunctor) (Extender, error)
	// AsLocation(*conjunctor.Conjunctor) (Extender, error)
	// AsApproach(*conjunctor.Conjunctor) (Extender, error)
	// AsStart(*conjunctor.Conjunctor) (Extender, error)
	// AsEnd(*conjunctor.Conjunctor) (Extender, error)

	Asserted(*conjunctor.Conjunctor) (Extender, error)

	Negated(*conjunctor.Conjunctor) (Extender, error)
	// Potentially(*conjunctor.Conjunctor) (Extender, error)
	// Politely(*conjunctor.Conjunctor) (Extender, error)
	// Desirably(*conjunctor.Conjunctor) (Extender, error)
	// Purposefully(*conjunctor.Conjunctor) (Extender, error)
	// Attributing(*conjunctor.Conjunctor) (Extender, error)
	// Terminate(*conjunctor.Conjunctor) (Extender, error)
	// Pause(*conjunctor.Conjunctor) (Extender, error)
	// Statified(*conjunctor.Conjunctor) (Extender, error)
	// Completed(*conjunctor.Conjunctor) (Extender, error)
	// Imperatively(*conjunctor.Conjunctor) (Extender, error)
	// Volitionally(*conjunctor.Conjunctor) (Extender, error)
	// Forbade(*conjunctor.Conjunctor) (Extender, error)
}

type UnimplementedExtender struct{}

func (u UnimplementedExtender) Asserted(*conjunctor.Conjunctor) (Extender, error) {
	return nil, fmt.Errorf("Unimplemented: Asserted()")
}

func (u UnimplementedExtender) Negated(*conjunctor.Conjunctor) (Extender, error) {
	return nil, fmt.Errorf("Unimplemented: Negated()")
}
