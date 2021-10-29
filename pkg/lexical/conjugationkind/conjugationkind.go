package conjugationkind

type ConjugationKind string

const (
	Unknown      ConjugationKind = "Unknown"
	Imperfective ConjugationKind = "Imperfective"
	Conjunctive  ConjugationKind = "Conjunctive"
	Attributive  ConjugationKind = "Attributive"
	Terminal     ConjugationKind = "Terminal"
	Volitional   ConjugationKind = "Volitional"
	Imperative   ConjugationKind = "Imperative"
	Conditional  ConjugationKind = "Conditional"
)
