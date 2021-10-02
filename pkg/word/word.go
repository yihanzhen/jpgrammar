package word

type Word struct {
	ganas        []rune
	kanjis       []rune
	writeOptions []WriteOption
}

type WriteOption int

const (
	HiraganaOnly WriteOption = iota
)

func NewWord(ganas, kanjis string, opts ...WriteOption) Word {
	return Word{
		ganas:        []rune(ganas),
		kanjis:       []rune(kanjis),
		writeOptions: opts,
	}
}

func (w Word) Write() string {
	for _, opt := range w.writeOptions {
		if opt == HiraganaOnly {
			return string(w.ganas)
		}
	}
	if len(w.kanjis) == 0 {
		return string(w.ganas)
	}
	return string(w.kanjis)
}
