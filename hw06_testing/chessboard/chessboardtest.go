package chessboard

import "testing"

func TestGenerateChessboard(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		want    string
		wantErr bool
	}{
		{
			name:    "Standard 5x5 board",
			size:    5,
			want:    "#   #   # \n  #   #   \n#   #   # \n  #   #   \n#   #   # \n",
			wantErr: false,
		},
		{
			name:    "Empty board (size 0)",
			size:    0,
			want:    "",
			wantErr: true,
		},
		{
			name:    "Single-cell board (size 1)",
			size:    1,
			want:    "# \n",
			wantErr: false,
		},
		{
			name:    "Negative size",
			size:    -5,
			want:    "",
			wantErr: true,
		},
		{
			name:    "Standard 2x2 board",
			size:    2,
			want:    "#   \n  # \n",
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GenerateChessboard(test.size)

			if (err != nil) != test.wantErr {
				t.Errorf("unexpected error state: got %v, wantErr %v", err, test.wantErr)
			}

			if got != test.want {
				t.Errorf("for size %d, expected:\n%q\ngot:\n%q", test.size, test.want, got)
			}
		})
	}
}
