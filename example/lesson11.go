package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/casekind"
)

// L11E1: 会議室にテーブルが7つあります。
func L11E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("会議室")
	b.Vocab.AddNoun("テーブル")
	b.Vocab.AddAdverb("7つ")
	b.Vocab.AddVerb("ある", "ある")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("会議室").As(casekind.Location)
	b.Append("テーブル").As(casekind.Subject)
	b.Append("7つ")
	b.Append("ある").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

// L11E2: 私は日本に1年います。
func L11E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("日本")
	b.Vocab.AddAdverb("1年")
	b.Vocab.AddVerb("いる", "いる")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("日本").As(casekind.Location)
	b.Append("1年")
	b.Append("いる").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}
