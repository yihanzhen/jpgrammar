package example

import "testing"

func TestLesson3(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L3E1,
			want:  "ここは食堂です",
		},
		{
			name:  "example 2",
			build: L3E2,
			want:  "電話はあそこです",
		},
		{
			name:  "example 3",
			build: L3E3,
			want:  "これはどこのカメラですか",
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
