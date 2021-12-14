package example

import "testing"

func TestLesson10(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L10E1,
			want:  "あそこに佐藤さんがいます",
		},
		{
			name:  "example 2",
			build: L10E2,
			want:  "机の上に写真があります",
		},
		{
			name:  "example 3",
			build: L10E3,
			want:  "家族はニューヨークにいます",
		},
		{
			name:  "example 4",
			build: L10E4,
			want:  "東京ディズニーランドは千葉県にあります",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.build()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
