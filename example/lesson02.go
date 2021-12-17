package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/builder/vocabulary/word"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
)

// L2E1: これは辞書です。
func L2E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("これ")
	b.Vocab.AddNoun("辞書")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("これ").Mark(particle.Topic)
	b.Append("辞書").Asserted()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L2E2: これはコンピュータの本です。
func L2E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("これ")
	b.Vocab.AddNoun("コンピュータ")
	b.Vocab.AddNoun("本")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("これ").Mark(particle.Topic)
	b.Append("コンピュータ").Attributing("本").Asserted()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L2E3: それはわたしの傘です。
func L2E3() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("それ")
	b.Vocab.AddNoun("わたし")
	b.Vocab.AddNoun("傘")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("それ").Mark(particle.Topic)
	b.Append("わたし").Attributing("傘").Asserted()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L2E4: この傘はわたしのです。
func L2E4() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("この傘")
	b.Vocab.AddNoun("わたし")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("この傘").Mark(particle.Topic)
	b.Append("わたし").Attributing(word.OmittedNoun).Asserted()

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}
