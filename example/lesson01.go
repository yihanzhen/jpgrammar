package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
)

// L1E1: わたしはマイク・ミラーです。
func L1E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("わたし")
	b.Vocab.AddNoun("マイク・ミラー")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("わたし").Mark(particle.Topic)
	b.Append("マイク・ミラー").Asserted()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L1E2: サントスさんは学生ではありません。
func L1E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("サントスさん")
	b.Vocab.AddNoun("学生")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("サントスさん").Mark(particle.Topic)
	b.Append("学生").Asserted().Negated()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L1E3: ミラーさんは会社員ですか。
func L1E3() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("ミラーさん")
	b.Vocab.AddNoun("会社員")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("ミラーさん").Mark(particle.Topic)
	b.Append("会社員").Asserted()
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

// L1E4: サントスさんも会社員ですか。
func L1E4() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("サントスさん")
	b.Vocab.AddNoun("会社員")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("サントスさん").Mark(particle.LikewiseTopic)
	b.Append("会社員").Asserted()
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
