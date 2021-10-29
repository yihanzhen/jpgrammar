package word

import "testing"

func TestChangeLastRuneToCol(t *testing.T) {

	cases := []struct {
		input Word
		want  Word
		col   int
	}{
		{
			input: MustWord("飲む", "のむ"),
			want:  MustWord("飲み", "のみ"),
			col:   1,
		},
	}

	for _, tc := range cases {
		got, err := tc.input.ChangeLastRuneTo(ToCol(1))
		if err != nil {
			t.Fatalf("got %v, want success", err)
		}
		if got != tc.want {
			t.Fatalf("got %v, want %v", got, tc.want)
		}
	}
}
