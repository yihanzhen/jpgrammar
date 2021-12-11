package example

import "testing"

func TestLesson8(t *testing.T) {

	cases := []struct {
		name  string
		want  string
		build func() (string, error)
	}{
		{
			name:  "example 1",
			build: L8E1,
			want:  "櫻はきれいです",
		},
		{
			name:  "example 2",
			build: L8E2,
			want:  "富士山は高いです",
		},
		{
			name:  "example 3",
			build: L8E3,
			want:  "櫻はきれいな花です",
		},
		{
			name:  "example 4",
			build: L8E4,
			want:  "富士山は高い山です",
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
