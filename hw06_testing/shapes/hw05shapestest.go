package shapes

import (
	"errors"
	"testing"
)

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("нет фигуры")
	}
	return shape.Area(), nil
}

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name    string
		shape   any
		want    float64
		wantErr bool
	}{
		{
			name:    "Rectangle",
			shape:   Rectangle{Weight: 5, Height: 10},
			want:    50.0,
			wantErr: false,
		},
		{
			name:    "Circle",
			shape:   Circle{Radius: 10},
			want:    314,
			wantErr: false,
		},
		{
			name:    "Rectangle with zero width",
			shape:   Rectangle{Weight: 0, Height: 10},
			want:    0,
			wantErr: false,
		},
		{
			name:    "Invalid type (not a shape)",
			shape:   "not a shape",
			want:    0,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := calculateArea(test.shape)
			if (err != nil) != test.wantErr {
				t.Errorf("unexpected error state for %s: got error = %v, wantErr = %v", test.name, err, test.wantErr)
			}

			if result != test.want {
				t.Errorf("for %s, expected %.2f, got %.2f", test.name, test.want, result)
			}
		})
	}
}
