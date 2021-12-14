package example

import "testing"

func TestLesson9(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L9E1,
			want:  "私はイタリア料理が好きです",
		},
		{
			name:  "example 2",
			build: L9E2,
			want:  "私は日本語が少しわかります",
		},
		{
			name:  "example 3",
			build: L9E3,
			want:  "今日は子供の誕生日ですから早く帰ります",
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
