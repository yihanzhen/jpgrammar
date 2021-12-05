package wordkind

// WordKind represents the kind of a word.
type WordKind int

const (
	Unknown WordKind = iota

	// Verb represents a verb.
	Verb

	// Noun represents a noun.
	Noun

	// Adjective represents an adjective, or i-adjective.
	Adjective

	// AdjNoun represents an adjectival noun, or na-adjective.
	AdjNoun

	// Adverb represents an adverb.
	Adverb

	// AuxVerb represents an aux verb.
	AuxVerb

	// Parcicle represents a particle.
	Particle
)

// String implements the Stringer interface.
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
