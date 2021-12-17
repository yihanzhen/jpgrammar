package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
)

// L8E1: 櫻はきれいです。
func L8E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("櫻")
	b.Vocab.AddAdjNoun("きれい")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("櫻").Mark(particle.Topic)
	b.Append("きれい").Asserted()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

// L8E2: 富士山は高いです。
func L8E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("富士山")
	b.Vocab.AddAdjective("高い")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("富士山").Mark(particle.Topic)
	b.Append("高い").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

// L8E3: 櫻はきれいな花です。
func L8E3() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("櫻")
	b.Vocab.AddAdjNoun("きれい")
	b.Vocab.AddNoun("花")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("櫻").Mark(particle.Topic)
	b.Append("きれい").Attributing("花").Asserted()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

//L8E4: 富士山は高い山です。
func L8E4() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("富士山")
	b.Vocab.AddAdjective("高い")
	b.Vocab.AddNoun("山")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("富士山").Mark(particle.Topic)
	b.Append("高い").Attributing("山").Asserted()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}
