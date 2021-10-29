package extender

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
)

type Extender interface {

	// To assert the component to be extended. The component
	// should be a noun, an adjectival noun, or anything
	// that functions like a nominal.
	Asserted(*conjunctor.Conjunctor) (Extender, error)

	// To negate the component to be extended. The component
	// should be a verb, an adjective, certain aux verbs such
	// as です, だ.
	Negated(*conjunctor.Conjunctor) (Extender, error)

	// AsSubject(*conjunctor.Conjunctor) (Extender, error)
	// AsObject(*conjunctor.Conjunctor) (Extender, error)
	// AsDirection(*conjunctor.Conjunctor) (Extender, error)
	// AsVenue(*conjunctor.Conjunctor) (Extender, error)
	// AsTimestamp(*conjunctor.Conjunctor) (Extender, error)
	// AsLocation(*conjunctor.Conjunctor) (Extender, error)
	// AsApproach(*conjunctor.Conjunctor) (Extender, error)
	// AsStart(*conjunctor.Conjunctor) (Extender, error)
	// AsEnd(*conjunctor.Conjunctor) (Extender, error)
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

// UnimplementedExtender is a default implementation of the Extender interface.
type UnimplementedExtender struct {
	Name string
}

func (u UnimplementedExtender) Asserted(*conjunctor.Conjunctor) (Extender, error) {
	return nil, fmt.Errorf("%sAsserted()", u.getName())
}

func (u UnimplementedExtender) Negated(*conjunctor.Conjunctor) (Extender, error) {
	return nil, fmt.Errorf("%sNegated()", u.getName())
}

func (u UnimplementedExtender) getName() string {
	if u.Name != "" {
		return u.Name + ": "
	}
	return ""
}
