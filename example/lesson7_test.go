package example

import "testing"

func TestLesson7(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L7E1,
			want:  "私はワープロで手紙を書きます",
		},
		{
			name:  "example 2",
			build: L7E2,
			want:  "私は木村さんに花をあげます",
		},
		{
			name:  "example 3",
			build: L7E3,
			want:  "私はカリナさんにチョコレートをもらいます",
		},
		{
			name:  "example 4",
			build: L7E4,
			want:  "もう新幹線の切符を買いましたか",
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
