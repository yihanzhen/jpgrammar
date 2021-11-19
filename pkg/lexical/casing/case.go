package casing

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
)
