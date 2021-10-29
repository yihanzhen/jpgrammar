package word

import "testing"

func TestCheckLastRuneCol(t *testing.T) {
	cases := []struct {
		input Word
		col   int
	}{
		{
			input: MustWord("食べる", "たべる"),
			col:   2,
		},
	}
	for _, tc := range cases {
		if got := tc.input.CheckLastRuneCol(tc.col); got != true {
			t.Errorf("CheckLastRuneCol(%v): got %v, want %v", tc.col, got, true)
		}
	}
}
