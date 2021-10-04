package word

import (
	"fmt"
	"strings"

	"github.com/yihanzhen/jpgrammar/pkg/kana"
)

type Word struct {
	canonical string
	display   string
}

func NewWord(canonical, display string) (Word, error) {
	if display == "" {
		display = canonical
	}
	w := Word{
		canonical: canonical,
		display:   display,
	}
	if err := Check(w, CanonicalNotEmpty, SuffixHiraganaOnly); err != nil {
		return Word{}, err
	}
	return w, nil
}

func (w Word) Write() string {
	return w.display
}

type CheckCondtion func(word Word) error

func CanonicalNotEmpty(word Word) error {
	if len(word.canonical) == 0 {
		return fmt.Errorf("kanas of the word is empty")
	}
	return nil
}

func SuffixNotEmpty(word Word) error {
	if len(word.suffix) == 0 {
		return fmt.Errorf("kanas of the word is empty")
	}
	return nil
}

func SuffixHiraganaOnly(word Word) error {
	for _, r := range word.suffix {
		if !kana.IsHiragana(r) {
			return fmt.Errorf("suffix has non-hiragana rune: %v", r)
		}
	}
	return nil
}

func SuffixLastRuneIs(word Word, r rune) func(word Word) error {
	return func(word Word) error {
		if err := SuffixNotEmpty(word); err != nil {
			return err
		}
		sr := []rune(word.suffix)
		lr := sr[len(sr)-1]
		if lr != r {
			return fmt.Errorf("SuffixLastRuneIs: got %v, want %v", lr, r)
		}
		return nil
	}
}

func SuffixLastRuneIsCol(word Word, col int) func(word Word) error {
	return func(word Word) error {
		if err := SuffixNotEmpty(word); err != nil {
			return err
		}
		sr := []rune(word.suffix)
		lr := sr[len(sr)-1]
		pos, ok := kana.HiraganaMap[lr]
		if !ok {
			return fmt.Errorf("SuffixLastRuneIsCol, invalid last rune: got %v, want hiragana", lr)
		}
		if pos[1] != col {
			return fmt.Errorf("SuffixLastRuneIsCol, invalid col: got %v, want %v", pos[1], col)
		}
		return nil
	}
}

func Check(word Word, conds ...CheckCondtion) error {
	for _, cond := range conds {
		if err := cond(word); err != nil {
			return err
		}
	}
	return nil
}

func LastKanaToCol(word Word, col int) (Word, error) {
	newSuf, err := lastKanaToCol(word.suffix, col)
	if err != nil {
		return Word{}, err
	}
	word.suffix = newSuf
	return word, nil
}

func LastKanaAToWa(word Word) Word {
	if len(word.suffix) > 0 && []rune(word.suffix)[len(word.suffix)-1] == 'あ' {
		suf := []rune(word.suffix)
		suf[len(word.suffix)-1] = 'わ'
		word.suffix = string(suf)
	}
	return word
}

func TrimLastKana(word Word) (Word, error) {
	if word.suffix == "" {
		return Word{}, fmt.Errorf("suffix must not be empty")
	}
	rs := []rune(word.suffix)
	word.suffix = string(rs[0 : len(rs)-1])
	return word, nil
}

func AppendSuffix(word Word, suf string) Word {
	word.suffix = word.suffix + suf
	return word
}

func (w Word) EndsWith(suf string) bool {
	return strings.HasSuffix(w.suffix, suf)
}

func lastKanaToCol(suf string, col int) (string, error) {
	if suf == "" {
		return "", fmt.Errorf("suffix must not be empty")
	}
	rs := []rune(suf)
	r := rs[len(rs)-1]
	r2, err := kana.Col(r, col)
	if err != nil {
		return "", err
	}
	rs[len(rs)-1] = r2
	return string(rs), nil
}
