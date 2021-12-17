package example

import (
	"fmt"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/types/casekind"
)

// L12E1: 昨日は雨でした。
func L12E1() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("昨日")
	b.Vocab.AddNoun("雨")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("昨日").Mark(particle.Topic)
	b.Append("雨").Asserted().Completed()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

// L12E2: 昨日は寒かったです。
func L12E2() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("昨日")
	b.Vocab.AddAdjective("寒い")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("昨日").Mark(particle.Topic)
	b.Append("寒い").Completed().Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

// L12E3: 北海道は九州より大きいです。
func L12E3() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("北海道")
	b.Vocab.AddNoun("九州")
	b.Vocab.AddAdjective("大きい")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("北海道").Mark(particle.Topic)
	b.Append("九州").As(casekind.ComparisonBase)
	b.Append("大きい").Politely()
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

// L12E4: 私は一年で夏が一番好きです。
func L12E4() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("私")
	b.Vocab.AddNoun("一年")
	b.Vocab.AddNoun("夏")
	b.Vocab.AddAdverb("一番")
	b.Vocab.AddAdjNoun("好き")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("私").Mark(particle.Topic)
	b.Append("一年").As(casekind.Range)
	b.Append("夏").As(casekind.Subject)
	b.Append("一番")
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

// L12E5: サッカーと野球とどちらが面白いですか。
func L12E5() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("サッカー")
	b.Vocab.AddNoun("野球")
	b.Vocab.AddNoun("どちら")
	b.Vocab.AddAdjective("面白い")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("サッカー").Mark(particle.CompleteList)
	b.Append("野球").Mark(particle.CompleteList)
	b.Append("どちら").As(casekind.Subject)
	b.Append("面白い").Politely()
	b.Mark(particle.Uncertainty)
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}

// L12E6: スポーツで何が一番面白いですか。
func L12E6() (string, error) {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("スポーツ")
	b.Vocab.AddNoun("何")
	b.Vocab.AddAdverb("一番")
	b.Vocab.AddAdjective("面白い")
	if b.Vocab.GetError() != nil {
		return "", fmt.Errorf("Adding vocabulary has errors: %v", b.Vocab.GetError())
	}

	b.Append("スポーツ").As(casekind.Range)
	b.Append("何").As(casekind.Subject)
	b.Append("一番")
	b.Append("面白い").Politely()
	b.Mark(particle.Uncertainty)
	if b.Diag.HasErrors() {
		return "", fmt.Errorf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		return "", fmt.Errorf("Printing sentence has errors: %v", err)
	}
	return sentence, nil
}
