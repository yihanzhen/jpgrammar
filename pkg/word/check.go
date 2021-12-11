package word

import (
	"strings"

	"github.com/yihanzhen/jpgrammar/pkg/kana"
)

// CheckSuffix checks if the conjugateRef of a word has a certain suffix.
func (w Word) CheckSuffix(suf string) bool {
	return strings.HasSuffix(w.conjugateRef, suf) || strings.HasSuffix(w.writing, suf)
}

// CheckLastRune checks if the conjugateRef of a word ends in a certain rune.
// Returns false if conjugateRef is unset.
func (w Word) CheckLastRune(r rune) bool {
	if w.conjugateRef == "" {
		return false
	}
	sr := []rune(w.conjugateRef)
	lr := sr[len(sr)-1]
	return r == lr
}

// CheckLastRune checks if the conjugateRef of a word ends in a certain column.
// Returns false if conjugateRef is unset.
func (w Word) CheckLastRuneCol(col int) bool {
	sr := []rune(w.conjugateRef)
	lr := sr[len(sr)-1]
	pos, ok := kana.HiraganaMap[lr]
	return ok && pos[1] == col
}

// CheckRune checks if a certain rune of the conjugateRef of a word meets a certain requirement.
func (w Word) CheckRune(getter func(w Word) rune, checker func(r rune) bool) bool {
	return checker(getter(w))
}

// NthLastRune returns the function that returns the nth last rune of the conjugateRef of a word.
func NthLastRune(n int) func(w Word) rune {
	return func(w Word) rune {
		return w.NthLastRune(n)
	}
}

// NthLastRune returns the function that returns the nth last rune of the conjugateRef of a word.
func (w Word) NthLastRune(n int) rune {
	sr := []rune(w.conjugateRef)
	return sr[len(sr)-n-1]
}

// IsCol returns the function that returns whether a rune is a hiragana of a certain column.
func IsCol(col int) func(r rune) bool {
	return func(r rune) bool {
		return kana.IsCol(r, col)
	}
}
