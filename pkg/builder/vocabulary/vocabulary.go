package vocabulary

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder/conjunctor"
	"github.com/yihanzhen/jpgrammar/pkg/builder/extender"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/verb"
)

type Vocabulary struct {
	dict   map[string]Word
	errors []error
}

func NewVocabulary() *Vocabulary {
	return &Vocabulary{
		dict: map[string]Word{},
	}
}

type Word interface {
	conjunctor.Conjunctable
	extender.Extender
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

func (v *Vocabulary) AddNoun(canonical, display string) {

}

func (v *Vocabulary) GetWord(str string) (Word, error) {
	w, ok := v.dict[str]
	if !ok {
		return nil, fmt.Errorf("GetWord: word %s does not exist", str)
	}
	return w, nil
}

func (v *Vocabulary) GetError() error {
	return fmt.Errorf("Vocabulary has errors: %v", v.errors)
}
