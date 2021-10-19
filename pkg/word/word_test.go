package word

import "testing"

func TestNewWord(t *testing.T) {
	cases := []struct {
		canonical string
		display   string
		wantErr   bool
	}{
		{
			canonical: "たべる",
			display:   "食べる",
		},
		{
			canonical: "食べる",
			display:   "食べる",
			wantErr:   true,
		},
	}
	for _, tc := range cases {
		_, err := NewWord(tc.canonical, tc.display)
		if (err != nil) != tc.wantErr {
			t.Fatalf("NewWord: got %v, want %v", err, tc.wantErr)
		}
	}
}
