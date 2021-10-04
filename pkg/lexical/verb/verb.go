package verb

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Verb struct {
	word.Word
	forceGodanConj bool
}

type NewVerbOption int

const (
	ForceGodanConjugation NewVerbOption = iota
	WriteHiraganaOnly
)

func NewVerb(kanas, kanjis string, opts ...NewVerbOption) (Verb, error) {
	if kanas == "" {
		return Verb{}, fmt.Errorf("input kanas can't be empty")
	}
	lr := lastRune(kanas)
	if lk := lastRune(kanjis); lr != lk {
		return Verb{}, fmt.Errorf("last rune of kanas and kanjis must be the same, got %v and %v", lr, lk)
	}

	kanas = trimLastRune(kanas)
	kanjis = trimLastRune(kanjis)

	var writeOpts []word.WriteOption
	var forceGodanConj bool
	for _, opt := range opts {
		if opt == WriteHiraganaOnly {
			writeOpts = append(writeOpts, word.HiraganaOnly)
		}
		if opt == ForceGodanConjugation {
			forceGodanConj = true
		}
	}

	w, err := word.NewWord(kanas, kanjis, string([]rune{lr}), writeOpts...)
	if err != nil {
		return Verb{}, fmt.Errorf("NewVerb: %v", err)
	}
	v := Verb{
		Word:           w,
		forceGodanConj: forceGodanConj,
	}
	return v, nil
}

func (v Verb) GetConjugator() VerbConjugator {
	if v.forceGodanConj {
		return &GodanVerbConjugator{
			verb: v,
		}
	}
	if v.EndsWith("る") && v.
	return &GodanVerbConjugator{
		verb: v,
	}
}

func lastRune(str string) rune {
	rs := []rune(str)
	return rs[len(rs)-1]
}

func trimLastRune(str string) string {
	rs := []rune(str)
	return string(rs[0 : len(rs)-1])
}
