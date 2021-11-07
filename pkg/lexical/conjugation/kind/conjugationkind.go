package kind

// ConjugationKind represents a conjugation kind.
type ConjugationKind int

const (
	// Unknown is the default value.
	Unknown ConjugationKind = iota

	// Irrealis is the imperfective conjugation, or 未然形.
	// For verbs, this conjugation form is used to conjunct passive/potential auxverb "られる",
	// negative auxverb "ない", causative auxverb "させる".

	// (TODO: find source) Many sources consider volitional forms to be derived from the Irrelis
	// conjugation; although in modern Japanese the volitional form has its own conjugation table.
	Irrealis

	// Conjunctive is the conjunctive conjugation, or 連用形.
	// For verbs, this conjugation form is used to conjunct polite auxverb "ます", desirable
	// auxverb "たい", purpose particle "に", other verbs to form compound verbs, or as a noun
	// without conjuncting anything.
	//
	// For adjectives, this conjugation form is used to conjunct negative auxverb "ない", perfective
	// auxverb "た".
	Conjunctive

	// Attributive is the attributive conjugation, or 連体形.
	// For verbs, adjectives, nouns and adjectival nouns, this conjugation is used to conjunct
	// nouns.
	Attributive
	Terminal
	Volitional
	Imperative
	Conditional
)

func (c ConjugationKind) String() string {
	switch c {
	case Irrealis:
		return "Irrealis"
	case Conjunctive:
		return "Conjunctive"
	case Attributive:
		return "Attributive"
	case Terminal:
		return "Terminal"
	case Volitional:
		return "Volitional"
	case Imperative:
		return "Imperative"
	case Conditional:
		return "Conditional"
	}
	return "Unknown"
}
