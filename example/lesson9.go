package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/casekind"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/verb"
)

// L9E1: 私はイタリア料理が好きです。
func L9E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("イタリア料理")
	b.Vocab.AddAdjNoun("好き")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("イタリア料理").As(casekind.Subject)
	b.Append("好き").Asserted()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

// L9E2: 私は日本語が少しわかります。
func L9E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("日本語")
	b.Vocab.AddAdverb("少し")
	b.Vocab.AddVerb("わかる", "わかる")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("日本語").As(casekind.Subject)
	b.Append("少し")
	b.Append("わかる").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

// L9E3: 今日は子供の誕生日ですから、早く帰ります。
func L9E3() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("今日")
	b.Vocab.AddNoun("子供")
	b.Vocab.AddNoun("誕生日")
	b.Vocab.AddAdverb("早く")
	b.Vocab.AddVerb("帰る", "かえる", verb.ForceTypeOneConjugation)
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("今日").Mark(particle.Topic)
	b.Append("子供").Attributing("誕生日").Asserted()
	b.Mark(particle.Reason)
	b.Append("早く")
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
