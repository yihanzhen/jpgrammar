package example

import "testing"

func TestLesson6(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L6E1,
			want:  "私はジュースを飲みます",
		},
		{
			name:  "example 2",
			build: L6E2,
			want:  "私は駅で新聞を買います",
		},
		{
			name:  "example 3",
			build: L6E3,
			want:  "一緒に神戸へ行きませんか",
		},
		{
			name:  "example 4",
			build: L6E4,
			want:  "ちょっと休みましょう",
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
