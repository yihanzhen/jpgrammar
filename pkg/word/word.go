package word

import "github.com/yihanzhen/jpgrammar/pkg/kana"

type Word struct {
	kanas        []rune
	kanjis       []rune
	writeOptions []WriteOption
}

type WriteOption int

const (
	HiraganaOnly WriteOption = iota
)

func NewWord(kanas, kanjis string, opts ...WriteOption) Word {
	return Word{
		kanas:        []rune(kanas),
		kanjis:       []rune(kanjis),
		writeOptions: opts,
	}
}

func (w Word) Write() string {
	for _, opt := range w.writeOptions {
		if opt == HiraganaOnly {
			return string(w.kanas)
		}
	}
	if len(w.kanjis) == 0 {
		return string(w.kanas)
	}
	return string(w.kanjis)
}

func LastKanaToCol(word Word, col int) (Word, error) {
	w := NewWord("", "", word.writeOptions...)
	copy(w.kanas, word.kanas)
	copy(w.kanjis, word.kanjis)

	r, err := kana.Col(w.kanas[len(w.kanas)-1], col)
	if err != nil {

	}
	w.kanas[len(w.kanas)-1] = r
	w.kanjis[len(w.kanjis)-1] = r
	return w, nil
}
