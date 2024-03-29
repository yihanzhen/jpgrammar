package vocabulary

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/vocabulary/word"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/adjective"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/adjnoun"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/adverb"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/noun"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/verb"
)

type Vocabulary struct {
	dict   map[string]word.Word
	errors []error
}

func NewVocabulary() *Vocabulary {
	return &Vocabulary{
		dict: map[string]word.Word{},
	}
}

func (v *Vocabulary) AddVerb(writing, conjRef string, opts ...verb.NewVerbOption) {
	vb, err := verb.NewVerb(writing, conjRef, opts...)
	if err != nil {
		v.errors = append(v.errors, fmt.Errorf("AddVerb: %v", err))
		return
	}
	if _, ok := v.dict[writing]; ok {
		v.errors = append(v.errors, fmt.Errorf("AddVerb: canonical %q already exists", writing))
		return
	}
	if _, ok := v.dict[conjRef]; ok {
		v.errors = append(v.errors, fmt.Errorf("AddVerb: display %q already exists", conjRef))
		return
	}
	v.dict[writing] = vb
	v.dict[conjRef] = vb
}

func (v *Vocabulary) AddNoun(writing string) {
	n, err := noun.NewNoun(writing)
	if err != nil {
		v.errors = append(v.errors, fmt.Errorf("AddNoun: %v", err))
		return
	}
	if _, ok := v.dict[writing]; ok {
		v.errors = append(v.errors, fmt.Errorf("AddNoun: writing %q already exists", writing))
		return
	}
	v.dict[writing] = n
}

func (v *Vocabulary) AddAdjNoun(writing string) {
	n, err := adjnoun.NewAdjNoun(writing)
	if err != nil {
		v.errors = append(v.errors, fmt.Errorf("AddAdjNoun: %v", err))
		return
	}
	if _, ok := v.dict[writing]; ok {
		v.errors = append(v.errors, fmt.Errorf("AddAdjNoun: writing %q already exists", writing))
		return
	}
	v.dict[writing] = n
}

func (v *Vocabulary) AddAdjective(writing string) {
	n, err := adjective.NewAdjective(writing)
	if err != nil {
		v.errors = append(v.errors, fmt.Errorf("AddAdjective: %v", err))
		return
	}
	if _, ok := v.dict[writing]; ok {
		v.errors = append(v.errors, fmt.Errorf("AddAdjective: writing %q already exists", writing))
		return
	}
	v.dict[writing] = n
}

func (v *Vocabulary) AddAdverb(writing string) {
	adv, err := adverb.NewAdverb(writing)
	if err != nil {
		v.errors = append(v.errors, fmt.Errorf("AddAdverb: %v", err))
		return
	}
	if _, ok := v.dict[writing]; ok {
		v.errors = append(v.errors, fmt.Errorf("AddAdverb: writing %q already exists", writing))
		return
	}
	v.dict[writing] = adv
}

func (v *Vocabulary) GetWord(str string) (word.Word, error) {
	if str == word.OmittedNoun {
		return noun.Omitted, nil
	}
	w, ok := v.dict[str]
	if !ok {
		return nil, fmt.Errorf("GetWord: word %s does not exist", str)
	}
	return w, nil
}

func (v *Vocabulary) GetError() error {
	if len(v.errors) == 0 {
		return nil
	}
	return fmt.Errorf("Vocabulary has errors: %v", v.errors)
}
