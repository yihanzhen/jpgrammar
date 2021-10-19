package extender

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
)

type Extender interface {
	AsSubject() Extender
	AsObject() Extender
	AsDirection() Extender
	AsVenue() Extender
	AsTimestamp() Extender
	AsLocation() Extender
	AsApproach() Extender
	AsStart() Extender
	AsEnd() Extender

	AsTopic() Extender
	Likewise() Extender

	Negated(*conjunctor.Conjunctor) (Extender, error)
	Potentially() Extender
	Politely() Extender
	Desirably() Extender
	Purposefully() Extender
	Attributing() Extender
	Terminate() Extender
	Pause() Extender
	Statified() Extender
	Completed() Extender
	Imperatively() Extender
	Volitionally() Extender
	Forbade() Extender
}

type UnimplementedExtender struct {
	SaveError func(err error)
	GetError  func() error
}

func (u UnimplementedExtender) AsSubject() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) AsObject() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}
func (u UnimplementedExtender) AsDirection() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) AsVenue() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) AsTimestamp() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) AsLocation() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) AsApproach() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) AsStart() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) AsEnd() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) AsTopic() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Likewise() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Negated(*conjunctor.Conjunctor) (Extender, error) {
	return nil, fmt.Errorf("Unimplemented: Negated()")
}

func (u UnimplementedExtender) Potentially() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}
func (u UnimplementedExtender) Politely() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Desirably() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Purposefully() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Attributing() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Terminate() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Pause() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Statified() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Completed() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Imperatively() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Volitionally() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}

func (u UnimplementedExtender) Forbade() Extender {
	if u.GetError() != nil {
		return u
	}
	u.SaveError(fmt.Errorf("Unimplemented: AsSubject()"))
	return u
}
