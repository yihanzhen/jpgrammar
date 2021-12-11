package adjnoun

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/conjugationkind"
	"github.com/yihanzhen/jpgrammar/pkg/word"
)

func (n AdjNoun) Conjugate(ck conjugationkind.ConjugationKind) (word.Word, error) {
	switch ck {
	case conjugationkind.Attributive:
		return n.Word.Append("な")
	}
	return word.Word{}, fmt.Errorf("AdjNoun.Conjugate: conjugationKind not Conjugatable: %v", ck)
}
