package example

import (
	"log"

	"github.com/yihanzhen/jpgrammar/pkg/builder"
	"github.com/yihanzhen/jpgrammar/pkg/lexical/particle"
)

func Example1() {
	b := builder.NewBuilder()
	b.Vocab.AddNoun("わたし", "わたし")
	b.Vocab.AddNoun("みらあ", "ミラー")

	if b.Diag.HasErrors() {
		log.Fatalf("Adding vocabulary has errors: %v", b.Diag.GetErrors())
	}
	b.Append("わたし")
	b.AppendParticle(particle.Topic)
	b.Append("ミラー").Asserted().Negated()

	if b.Diag.HasErrors() {
		log.Fatalf("Constructing sentence has errors: %v", b.Diag.GetErrors())
	}

	sentence, err := b.Build()
	if err != nil {
		log.Fatalf("Printing sentence has errors: %v", b.Diag.GetErrors())
	}
	log.Println(sentence)
}
