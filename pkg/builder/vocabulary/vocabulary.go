package vocabulary

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/vocabulary/word"
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

func (v *Vocabulary) AddVerb(canonical, display string, opts ...verb.NewVerbOption) {
	vb, err := verb.NewVerb(canonical, display, opts...)
	if err != nil {
		v.errors = append(v.errors, fmt.Errorf("AddVerb: %v", err))
		return
	}
	if _, ok := v.dict[canonical]; ok {
		v.errors = append(v.errors, fmt.Errorf("AddVerb: canonical %q already exists", canonical))
		return
	}
	if _, ok := v.dict[display]; ok {
		v.errors = append(v.errors, fmt.Errorf("AddVerb: display %q already exists", display))
		return
	}
	v.dict[canonical] = vb
	v.dict[display] = vb
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
	return fmt.Errorf("Vocabulary has errors: %v", v.errors)
}
