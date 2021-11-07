package noun

import (
	"testing"

	"github.com/yihanzhen/jpgrammar/pkg/lexical/conjugation/kind"
)

func TestConjugate(t *testing.T) {
	n, err := NewNoun("わたし")
	if err != nil {
		t.Fatalf("setup: %v", err)
	}
	attr, err := n.Conjugate(kind.Attributive)
	if err != nil {
		t.Errorf("Conjugate attributive: %v", err)
	}
	if attr.Write() != "わたしの" {
		t.Errorf("Conjugate result: got %q, want %q", attr.Write(), "わたしの")
	}
}
