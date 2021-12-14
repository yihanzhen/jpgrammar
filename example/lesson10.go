package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/casekind"
)

// L10E1: あそこに佐藤さんがいます。
func L10E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("あそこ")
	b.Vocab.AddNoun("佐藤さん")
	b.Vocab.AddVerb("いる", "いる")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("あそこ").As(casekind.Location)
	b.Append("佐藤さん").As(casekind.Subject)
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

// L10E2: 机の上に写真があります。
func L10E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("机")
	b.Vocab.AddNoun("上")
	b.Vocab.AddNoun("写真")
	b.Vocab.AddVerb("ある", "ある")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("机").Attributing("上").As(casekind.Location)
	b.Append("写真").As(casekind.Subject)
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

// L10E3: 家族はニューヨークにいます。
func L10E3() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("家族")
	b.Vocab.AddNoun("ニューヨーク")
	b.Vocab.AddVerb("いる", "いる")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("家族").Mark(particle.Topic)
	b.Append("ニューヨーク").As(casekind.Location)
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

//L10E4: 東京ディズニーランドは千葉県にあります。
func L10E4() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("東京ディズニーランド")
	b.Vocab.AddNoun("千葉県")
	b.Vocab.AddVerb("ある", "ある")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("東京ディズニーランド").Mark(particle.Topic)
	b.Append("千葉県").As(casekind.Location)
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
