package extender

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/casing"
)

type Extender interface {

	// Asserted is to assert the component to be extended. The component
	// should be a noun, an adjectival noun, or anything
	// that functions like a nominal.
	Asserted(*conjunctor.Conjunctor) (Extender, error)

	// Negated is to negate the component to be extended. The component
	// should be a verb, an adjective, or certain aux verbs such
	// as です, だ.
	Negated(*conjunctor.Conjunctor) (Extender, error)

	// Attributing is to make the component to be extended as attributives.
	// The component can be a verb, an adjective, a noun, an adjectival noun,
	// or certain aux verbs such as た.
	Attributing(conj *conjunctor.Conjunctor, attributed conjunctor.Conjunctable) (Extender, error)

	// As is to make the component to be extended as case. The component has to
	// be a noun.
	As(*conjunctor.Conjunctor, casing.CaseKind) (Extender, error)

	// Politely is to make the previous component polite. The component should either
	// be a verb, in which case calling this function will conjunct the polite aux verb
	// ます, an adjective, in which case calling this function will conjuct the polite
	// aux verb です, or the aux verb だ, in which case calling this function will replace
	// だ with the politer version です.
	Politely(*conjunctor.Conjunctor) (Extender, error)

	// Completed is to show the previous component has completed, or happened in the past.
	// The component can be a verb (た-form), an adjective (ka-tta form), negate aux verb nai,
	// aux verbs such as です, だ, ます and their respective negative forms.
	Completed(*conjunctor.Conjunctor) (Extender, error)

	// AsSubject(*conjunctor.Conjunctor) (Extender, error)
	// AsObject(*conjunctor.Conjunctor) (Extender, error)
	// AsDirection(*conjunctor.Conjunctor) (Extender, error)
	// AsVenue(*conjunctor.Conjunctor) (Extender, error)
	// AsLocation(*conjunctor.Conjunctor) (Extender, error)
	// AsApproach(*conjunctor.Conjunctor) (Extender, error)
	// Potentially(*conjunctor.Conjunctor) (Extender, error)
	// Desirably(*conjunctor.Conjunctor) (Extender, error)
	// Purposefully(*conjunctor.Conjunctor) (Extender, error)
	// Terminate(*conjunctor.Conjunctor) (Extender, error)
	// Pause(*conjunctor.Conjunctor) (Extender, error)
	// Statified(*conjunctor.Conjunctor) (Extender, error)
	// Imperatively(*conjunctor.Conjunctor) (Extender, error)
	// Volitionally(*conjunctor.Conjunctor) (Extender, error)
	// Forbade(*conjunctor.Conjunctor) (Extender, error)
}

// NewUnimplementedExtender creates a new UnimplementedExtender.
func NewUnimplementedExtender(name string) UnimplementedExtender {
	return UnimplementedExtender{
		Name: name,
	}
}

// UnimplementedExtender is a default implementation of the Extender interface.
type UnimplementedExtender struct {
	Name string
}

func (u UnimplementedExtender) Asserted(*conjunctor.Conjunctor) (Extender, error) {
	return u, fmt.Errorf("%sAsserted()", u.getName())
}

func (u UnimplementedExtender) Negated(*conjunctor.Conjunctor) (Extender, error) {
	return u, fmt.Errorf("%sNegated()", u.getName())
}

func (u UnimplementedExtender) Attributing(_ *conjunctor.Conjunctor, c conjunctor.Conjunctable) (Extender, error) {
	return u, fmt.Errorf("%s.Attributing(%s)", u.getName(), c)
}

func (u UnimplementedExtender) As(*conjunctor.Conjunctor, casing.CaseKind) (Extender, error) {
	return u, fmt.Errorf("%s.As()", u.getName())
}

func (u UnimplementedExtender) Politely(*conjunctor.Conjunctor) (Extender, error) {
	return u, fmt.Errorf("%s.Politely()", u.getName())
}

func (u UnimplementedExtender) Completed(*conjunctor.Conjunctor) (Extender, error) {
	return u, fmt.Errorf("%s.Completed()", u.getName())
}

func (u UnimplementedExtender) getName() string {
	if u.Name != "" {
		return u.Name + ": "
	}
	return ""
}
