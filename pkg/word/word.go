package word

import (
	"fmt"
	"log"

	"github.com/yihanzhen/jpgrammar/pkg/kana"
)

// Word is the representation of a word in a sentence.
type Word struct {
	// writing is the string representation of the word in written form.
	// It can consist of any unicode rune.
	writing string

	// conjugateRef is how the word is spelled in hiraganas.
	// It is used to determine what rules this word follow during conjugation.
	// Can be empty if the word doesn't conjugate (anything other than verbs
	// and adjectives).
	conjugateRef string
}

// Omitted represents an omitted word.
var Omitted Word = Word{
	writing:      "",
	conjugateRef: "",
}

// NewWord creates a new Word.
func NewWord(writing, conjRef string) (Word, error) {
	if writing == "" {
		return Word{}, fmt.Errorf("NewWord: writing cannot be empty")
	}
	if !kana.IsHiraganaStr(conjRef) {
		return Word{}, fmt.Errorf("NewWord: got conjRef %s, want hiraganas only", conjRef)
	}
	w := Word{
		writing:      writing,
		conjugateRef: conjRef,
	}
	return w, nil
}

// Write returns the string representation of the word.
func (w Word) Write() string {
	return w.writing
}

// MustWord calls NewWord, but panics if any error occurs.
func MustWord(writing, conjRef string) Word {
	w, err := NewWord(writing, conjRef)
	if err != nil {
		log.Fatalf("MustWord: unable to create new word: %v", err)
	}
	return w
}
