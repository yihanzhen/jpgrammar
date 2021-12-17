package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/casekind"
)

// L6E1: 私はジュースを飲みます。
func L6E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("ジュース")
	b.Vocab.AddVerb("飲む", "のむ")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("ジュース").As(casekind.Object)
	b.Append("飲む").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L6E2: 私は駅で新聞を買います。
func L6E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("駅")
	b.Vocab.AddNoun("新聞")
	b.Vocab.AddVerb("買う", "かう")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("駅").As(casekind.Venue)
	b.Append("新聞").As(casekind.Object)
	b.Append("買う").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L6E3: 一緒に神戸へ行きませんか。
func L6E3() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("一緒に")
	b.Vocab.AddNoun("神戸")
	b.Vocab.AddVerb("行く", "いく")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("一緒に")
	b.Append("神戸").As(casekind.Direction)
	b.Append("行く").Politely().Negated()
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

//L6E4: ちょっと休みましょう。
func L6E4() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddAdverb("ちょっと")
	b.Vocab.AddVerb("休む", "やすむ")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("ちょっと")
	b.Append("休む").Politely().Volitionally()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}
