package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/casekind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/verb"
)

// L5E1: 私は京都へ行きます。
func L5E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("京都")
	b.Vocab.AddVerb("行く", "いく")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("京都").As(casekind.Direction)
	b.Append("行く").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

// L5E2: 私はタクシーでうちへ帰ります。
func L5E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("タクシー")
	b.Vocab.AddNoun("うち")
	b.Vocab.AddVerb("帰る", "かえる", verb.ForceTypeOneConjugation)
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("タクシー").As(casekind.Instrument)
	b.Append("うち").As(casekind.Direction)
	b.Append("帰る").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

// L5E3: 私は家族と日本へ来ました。
func L5E3() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("家族")
	b.Vocab.AddNoun("日本")
	b.Vocab.AddVerb("来る", "くる")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("家族").As(casekind.Companion)
	b.Append("日本").As(casekind.Direction)
	b.Append("来る").Politely().Completed()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}
