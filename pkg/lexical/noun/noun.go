package noun

import (
	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/wordkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

type Noun struct {
	conjunctor.DefaultConjunctable
	word.Word
}

func (n Noun) GetWordKind() wordkind.WordKind {
	return wordkind.Noun
}
