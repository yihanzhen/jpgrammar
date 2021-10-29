package word

import "testing"

func TestNewWord(t *testing.T) {
	cases := []struct {
		conjRef string
		writing string
		wantErr bool
	}{
		{
			conjRef: "たべる",
			writing: "食べる",
		},
		{
			conjRef: "食べる",
			writing: "食べる",
			wantErr: true,
		},
	}
	for _, tc := range cases {
		_, err := NewWord(tc.writing, tc.conjRef)
		if (err != nil) != tc.wantErr {
			t.Fatalf("NewWord: got %v, want %v", err, tc.wantErr)
		}
	}
}
