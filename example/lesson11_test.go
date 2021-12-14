package example

import "testing"

func TestLesson11(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L11E1,
			want:  "会議室にテーブルが7つあります",
		},
		{
			name:  "example 2",
			build: L11E2,
			want:  "私は日本に1年います",
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
