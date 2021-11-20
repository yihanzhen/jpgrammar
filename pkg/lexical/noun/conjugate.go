package noun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

func (n Noun) Conjugate(ck conjugationkind.ConjugationKind) (word.Word, error) {
	switch ck {
	case conjugationkind.Attributive:
		return n.Word.Append("の")
	}
	return word.Word{}, fmt.Errorf("Verb.Conjugate: conjugationKind not Conjugatable: %v", ck)
}
