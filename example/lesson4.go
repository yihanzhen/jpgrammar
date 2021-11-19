package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/casing"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
)

// L4E1: 今、4時五分です。
func L4E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("今")
	b.Vocab.AddNoun("4時五分")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("今").As(casing.Time)
	b.Append("4時五分").Asserted()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L4E2: 私は9時から5時まで働きます。
func L4E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("9時")
	b.Vocab.AddNoun("5時")
	b.Vocab.AddVerb("働く", "はたらく")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("9時").As(casing.Start)
	b.Append("5時").As(casing.End)
	b.Append("働く").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L4E3: 私は朝6時に起きます。
func L4E3() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("朝6時")
	b.Vocab.AddVerb("起く", "おく")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("朝6時").As(casing.Timestamp)
	b.Append("起く").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

//L4E4: 私は昨日勉強しました。
func L4E4() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("昨日")
	b.Vocab.AddVerb("勉強する", "べんきょうする")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("昨日").As(casing.Time)
	b.Append("勉強する").Politely().Completed()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}
