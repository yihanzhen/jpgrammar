package word

import (
	"fmt"
	"log"

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

func MustWord(canonical, display string) Word {
	w, err := NewWord(canonical, display)
	if err != nil {
		log.Fatalf("mustWord: unable to create new word: %v", err)
	}
	return w
}
