package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
)

// L3E1: ここは食堂です。
func L3E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("ここ")
	b.Vocab.AddNoun("食堂")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("ここ").Mark(particle.Topic)
	b.Append("食堂").Asserted()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L2E2: 電話はあそこです。
func L3E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("電話")
	b.Vocab.AddNoun("あそこ")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("電話").Mark(particle.Topic)
	b.Append("あそこ").Asserted()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L3E3: これはどこのカメラですか。
func L3E3() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("これ")
	b.Vocab.AddNoun("どこ")
	b.Vocab.AddNoun("カメラ")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("これ").Mark(particle.Topic)
	b.Append("どこ").Attributing("カメラ").Asserted()
	b.Mark(particle.Uncertainty)
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}
