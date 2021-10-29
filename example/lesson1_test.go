package example

import "testing"

func TestLesson1(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L1E1,
			want:  "わたしはマイク・ミラーです",
		},
		{
			name:  "example 2",
			build: L1E2,
			want:  "サントスさんは学生ではありません",
		},
		{
			name:  "example 3",
			build: L1E3,
			want:  "ミラーさんは会社員ですか",
		},
		{
			name:  "example 4",
			build: L1E4,
			want:  "サントスさんも会社員ですか",
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
