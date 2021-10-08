package word

import (
	"strings"

	"github.com/yihanzhen/jpgrammar/pkg/kana"
)

func (w Word) CheckSuffix(suf string) bool {
	return strings.HasSuffix(w.canonical, suf)
}

func (w Word) CheckLastRune(r rune) bool {
	sr := []rune(w.canonical)
	lr := sr[len(sr)-1]
	return r == lr
}

func (w Word) CheckLastRuneCol(col int) bool {
	sr := []rune(w.canonical)
	lr := sr[len(sr)-1]
	pos, ok := kana.HiraganaMap[lr]
	return ok && pos[1] == col
}

func (w Word) CheckRune(getter func(w Word) rune, checker func(r rune) bool) bool {
	return checker(getter(w))
}

func NthLastRune(n int) func(w Word) rune {
	return func(w Word) rune {
		return w.NthLastRune(n)
	}
}

func (w Word) NthLastRune(n int) rune {
	sr := []rune(w.canonical)
	return sr[len(sr)-n-1]
}

func IsCol(col int) func(r rune) bool {
	return func(r rune) bool {
		return kana.IsCol(r, col)
	}
}
