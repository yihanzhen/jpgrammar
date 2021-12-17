package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/casekind"
)

// L7E1: 私はワープロで手紙を書きます。
func L7E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("ワープロ")
	b.Vocab.AddNoun("手紙")
	b.Vocab.AddVerb("書く", "かく")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("ワープロ").As(casekind.Instrument)
	b.Append("手紙").As(casekind.Object)
	b.Append("書く").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L7E2: 私は木村さんに花をあげます。
func L7E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("木村さん")
	b.Vocab.AddNoun("花")
	b.Vocab.AddVerb("あげる", "あげる")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("木村さん").As(casekind.IndirectObject)
	b.Append("花").As(casekind.Object)
	b.Append("あげる").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

// L7E3: 私はカリナさんにチョコレートをもらいます。
func L7E3() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("カリナさん")
	b.Vocab.AddNoun("チョコレート")
	b.Vocab.AddVerb("もらう", "もらう")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("カリナさん").As(casekind.IndirectObject)
	b.Append("チョコレート").As(casekind.Object)
	b.Append("もらう").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence, nil
}

//L7E4: もう新幹線の切符を買いましたか。
func L7E4() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddAdverb("もう")
	b.Vocab.AddNoun("新幹線")
	b.Vocab.AddNoun("切符")
	b.Vocab.AddVerb("買う", "かう")
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("もう")
	b.Append("新幹線").Attributing("切符").As(casekind.Object)
	b.Append("買う").Politely().Completed()
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
