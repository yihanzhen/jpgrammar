package word

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/kana"
)

// Append appends a suffix to the Word.
func (w Word) Append(str string) (Word, error) {
	var wt, conjRef string
	if w.conjugateRef != "" {
		conjRef = w.conjugateRef + str
	}
	wt = w.writing + str
	w2, err := NewWord(wt, conjRef)
	if err != nil {
		return Word{}, fmt.Errorf("Word.Append: %v", err)
	}
	return w2, nil
}

// TrimLastRune removes the last rune from the Word.
// It errors if the conjugateRef of the word is unset because
// there is no need to call this function for a unconjugatable word.
func (w Word) TrimLastRune() (Word, error) {
	if w.conjugateRef == "" {
		return Word{}, fmt.Errorf("Word.TrimLastRune: conjugateRef is unset")
	}
	wa := []rune(w.writing)
	wt := string(wa[0 : len(wa)-1])
	ca := []rune(w.conjugateRef)
	conjRef := string(ca[0 : len(ca)-1])
	w2, err := NewWord(wt, conjRef)
	if err != nil {
		return Word{}, fmt.Errorf("Word.TrimLastRune: %v", err)
	}
	return w2, nil
}

// ChangeLastRune changes the last rune of the Word.
// It errors if the conjugateRef of the word is unset because
// there is no need to call this function for a unconjugatable word.
func (w Word) ChangeLastRune(r rune) (Word, error) {
	if w.conjugateRef == "" {
		return Word{}, fmt.Errorf("Word.ChangeLastRune: conjugateRef is unset")
	}
	wa := []rune(w.writing)
	wt := string(append(wa[0:len(wa)-1], r))
	ca := []rune(w.conjugateRef)
	conjRef := string(append(ca[0:len(ca)-1], r))
	w2, err := NewWord(wt, conjRef)
	if err != nil {
		return Word{}, fmt.Errorf("Word.ChangeLastRune: %v", err)
	}
	return w2, nil
}

// ChangeLastRuneTo changes the last rune of the Word.
// It errors if the conjugateRef of the word is unset because
// there is no need to call this function for a unconjugatable word.
func (w Word) ChangeLastRuneTo(f func(r rune) (rune, error)) (Word, error) {
	if w.conjugateRef == "" {
		return Word{}, fmt.Errorf("Word.ChangeLastRuneTo: conjugateRef is unset")
	}
	wa := []rune(w.writing)
	wlr := wa[len(wa)-1]
	ca := []rune(w.conjugateRef)
	clr := ca[len(ca)-1]
	if wlr != clr {
		return Word{}, fmt.Errorf("Word.ChangeLastRuneTo: writing %v and conjugateRef %v have different suffix", w.writing, w.conjugateRef)
	}
	nr, err := f(clr)
	if err != nil {
		return Word{}, fmt.Errorf("ChangeLastRuneTo: %v", err)
	}
	wt := string(append(wa[0:len(wa)-1], nr))
	conjRef := string(append(ca[0:len(ca)-1], nr))
	w2, err := NewWord(wt, conjRef)
	if err != nil {
		return Word{}, fmt.Errorf("Word.ChangeLastRuneTo: %v", err)
	}
	return w2, nil
}

// ToCol returns a function that returns the hiragana of the same row as the input,
// but a different column.
func ToCol(col int) func(r rune) (rune, error) {
	return func(r rune) (rune, error) {
		r2, err := kana.Col(r, col)
		if err != nil {
			return ' ', fmt.Errorf("ToCol: %v", err)
		}
		return r2, nil
	}
}

// AToWa returns a function that returns `わ` if the input is `あ`.
func AToWa(r rune) (rune, error) {
	if r == 'あ' {
		return 'わ', nil
	}
	return r, nil
}
