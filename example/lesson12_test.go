package example

import "testing"

func TestLesson12(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L12E1,
			want:  "昨日は雨でした",
		},
		{
			name:  "example 2",
			build: L12E2,
			want:  "昨日は寒かったです",
		},
		{
			name:  "example 3",
			build: L12E3,
			want:  "北海道は九州より大きいです",
		},
		{
			name:  "example 4",
			build: L12E4,
			want:  "私は一年で夏が一番好きです",
		},
		{
			name:  "example 5",
			build: L12E5,
			want:  "サッカーと野球とどちらが面白いですか",
		},
		{
			name:  "example 6",
			build: L12E6,
			want:  "スポーツで何が一番面白いですか",
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
