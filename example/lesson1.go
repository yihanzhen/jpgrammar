package example

import (
	"log"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
)

// Example1: わたしはマイク・ミラーです。
func Example1() string {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("わたし")
	b.Vocab.AddNoun("マイク・ミラー")
	if b.Diag.HasErrors() {
		log.Fatalf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("わたし").Make(particle.Topic)
	b.Append("ミラー").Asserted()
	if b.Diag.HasErrors() {
		log.Fatalf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		log.Fatalf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence
}

// Example2: サントスさんは学生じゃありません。
func Example2() string {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("サントスさん")
	b.Vocab.AddNoun("学生")
	if b.Diag.HasErrors() {
		log.Fatalf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}

	b.Append("サントスさん").Make(particle.Topic)
	b.Append("学生").Asserted().Negated()
	if b.Diag.HasErrors() {
		log.Fatalf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		log.Fatalf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	return sentence
}

func Example3() string {
	return ""
}
