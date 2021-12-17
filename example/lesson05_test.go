package example

import "testing"

func TestLesson5(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L5E1,
			want:  "私は京都へ行きます",
		},
		{
			name:  "example 2",
			build: L5E2,
			want:  "私はタクシーでうちへ帰ります",
		},
		{
			name:  "example 3",
			build: L5E3,
			want:  "私は家族と日本へ来ました",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.build()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
