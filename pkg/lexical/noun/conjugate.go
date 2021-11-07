package noun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation/kind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

func (n Noun) Conjugate(ck kind.ConjugationKind) (word.Word, error) {
	switch ck {
	case kind.Attributive:
		return n.Word.Append("の")
	}
	return word.Word{}, fmt.Errorf("Verb.Conjugate: conjugationKind not Conjugatable: %v", ck)
}
