package casekind

type CaseKind int

const (
	Unknown CaseKind = iota

	// Start is the chronical or spatial starting point of a state or motion. It is marked by
	// the case marker から.
	Start

	// End is the chronical or spatial ending point of a state or motion. It is marked by
	// the case marker まで.
	// Note: a lot of sources do not consider まで a case marker, but in many entry-level usage
	// and examples there is little grammatical difference between から and まで. We will consider
	// まで a case marker until it doesn't work.
	End

	// Timestamp is the specific time when an event happens. It is marked by the case marker に.
	Timestamp

	// Time is the general time when an event happens. It does not need any case marker.
	Time

	// Direction is the direction of a motion. It is marked by the case marker へ.
	Direction

	// Instrument is the approach of an action. It is marked by the case marker で.
	Instrument

	// Companion is the companion that the action is done with. It is marked by the case marker と.
	Companion

	// Object is the object of an action. It is marked by the case marker を.
	Object

	// IndirectObject is the indirect object of an action, usually the person. It is marked by the case marker に.
	IndirectObject

	// Venue is the place the action takes place. It is marked by the case marker で.
	Venue

	// Subject is the subject of an action. It is marked by the case marker が.
	Subject

	// Location is the location of a state. It is marked by the case marker に.
	Location
)

// String implements the Stringer interface.
func (c CaseKind) String() string {
	switch c {
	case Start:
		return "start"
	case End:
		return "end"
	case Timestamp:
		return "timestamp"
	case Time:
		return "time"
	case Direction:
		return "direction"
	case Instrument:
		return "instrument"
	case Companion:
		return "companion"
	case Object:
		return "object"
	case IndirectObject:
		return "indirect object"
	case Venue:
		return "venue"
	case Subject:
		return "subject"
	case Location:
		return "location"
	}
	return "unknown"
}
