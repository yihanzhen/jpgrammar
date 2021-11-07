package example

import "testing"

func TestLesson2(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L2E1,
			want:  "これは辞書です",
		},
		{
			name:  "example 2",
			build: L2E2,
			want:  "これはコンピュータの本です",
		},
		{
			name:  "example 3",
			build: L2E3,
			want:  "それはわたしの傘です",
		},
		{
			name:  "example 4",
			build: L2E4,
			want:  "この傘はわたしのです",
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
