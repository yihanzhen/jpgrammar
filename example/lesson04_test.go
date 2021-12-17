package example

import "testing"

func TestLesson4(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L4E1,
			want:  "今4時五分です",
		},
		{
			name:  "example 2",
			build: L4E2,
			want:  "私は9時から5時まで働きます",
		},
		{
			name:  "example 3",
			build: L4E3,
			want:  "私は朝6時に起きます",
		},
		{
			name:  "example 4",
			build: L4E4,
			want:  "私は昨日勉強しました",
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
