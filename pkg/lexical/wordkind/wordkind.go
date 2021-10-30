package wordkind

type WordKind int

const (
	Unknown WordKind = iota
	Verb
	Noun
	Adjective
	AdjNoun
	AuxVerb
	Particle
)

func (wk WordKind) String() string {
	switch wk {
	case Verb:
		return "Verb"
	case Noun:
		return "Noun"
	case Adjective:
		return "Adjective"
	case AdjNoun:
		return "AdjNoun"
	case AuxVerb:
		return "AuxVerb"
	case Particle:
		return "Particle"
	}
	return "Unknown"
}
