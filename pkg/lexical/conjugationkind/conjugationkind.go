package conjugationkind

type ConjugationKind int

const (
	Unknown ConjugationKind = iota
	Imperfective
	Conjunctive
	Attributive
	Terminal
	Volitional
	Imperative
	Conditional
)
