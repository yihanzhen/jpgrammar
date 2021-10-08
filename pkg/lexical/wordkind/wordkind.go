package wordkind

type WordKind int

const (
	Unknown WordKind = iota
	Verb
	Noun
	Adjective
	AdjNoun
	Particle
)
