package word

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/kana"
)

func (w Word) Append(str string) (Word, error) {
	w2, err := NewWord(w.canonical+str, w.display+str)
	if err != nil {
		return Word{}, fmt.Errorf("Word.Append: %v", err)
	}
	return w2, nil
}

func (w Word) AppendWord(w2 Word) (Word, error) {
	w3, err := NewWord(w.canonical+w2.canonical, w.display+w2.display)
	if err != nil {
		return Word{}, fmt.Errorf("Word.AppendWord: %v", err)
	}
	return w3, nil
}

func (w Word) TrimLastRune() (Word, error) {
	c := []rune(w.canonical)
	d := []rune(w.display)
	w2, err := NewWord(string(c[0:len(c)-1]), string(d[0:len(d)-1]))
	if err != nil {
		return Word{}, fmt.Errorf("Word.TrimLastRune: %v", err)
	}
	return w2, nil
}

func (w Word) ChangeLastRune(r rune) (Word, error) {
	c := []rune(w.canonical)
	d := []rune(w.display)
	if c[len(c)-1] != d[len(d)-1] {
		return Word{}, fmt.Errorf("ChangeLastRune: canonial %s and display %s has different suffix", w.canonical, w.display)
	}
	w2, err := NewWord(string(append(c[0:len(c)-1], r)), string(append(d[0:len(d)-1], r)))
	if err != nil {
		return Word{}, fmt.Errorf("Word.ChangeLastRune: %v", err)
	}
	return w2, nil
}

func (w Word) ChangeLastRuneTo(f func(r rune) (rune, error)) (Word, error) {
	c := []rune(w.canonical)
	d := []rune(w.display)
	clr := c[len(c)-1]
	dlr := d[len(d)-1]
	if clr != dlr {
		return Word{}, fmt.Errorf("ChangeLastRuneTo: canonial %s and display %s has different suffix", w.canonical, w.display)
	}
	nr, err := f(clr)
	if err != nil {
		return Word{}, fmt.Errorf("ChangeLastRuneTo: %v", err)
	}
	w2, err := NewWord(string(append(c[0:len(c)-1], nr)), string(append(d[0:len(d)-1], nr)))
	if err != nil {
		return Word{}, fmt.Errorf("Word.ChangeLastRune: %v", err)
	}
	return w2, nil
}

func ToCol(col int) func(r rune) (rune, error) {
	return func(r rune) (rune, error) {
		r2, err := kana.Col(r, col)
		if err != nil {
			return ' ', fmt.Errorf("ToCol: %v", err)
		}
		return r2, nil
	}
}

func AToWa(r rune) (rune, error) {
	if r == 'あ' {
		return 'わ', nil
	}
	return r, nil
}
