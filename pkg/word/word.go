package word

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/kana"
)

type Word struct {
	canonical string
	display   string
}

func NewWord(canonical, display string) (Word, error) {
	if canonical == "" {
		return Word{}, fmt.Errorf("NewWord: canonical cannot be empty")
	}
	if !kana.IsHiraganaStr(canonical) {
		return Word{}, fmt.Errorf("NewWord: canonical must be hiraganas only")
	}
	if display == "" {
		display = canonical
	}
	w := Word{
		canonical: canonical,
		display:   display,
	}
	return w, nil
}

func (w Word) Write() string {
	return w.display
}
