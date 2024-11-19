package shapes

import "testing"

func Test_CalculateArea(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "Rectangle",
			shape: &Rectangle{Weight: 5, Height: 10},
			want:  50.0,
		},
		{
			name:  "Circle",
			shape: &Circle{Radius: 10},
			want:  314.16,
		},
		{
			name:  "Rectangle",
			shape: &Rectangle{Weight: 0, Height: 10},
			want:  0,
		},
		{
			name:  "Circle",
			shape: &Circle{Radius: -2},
			want:  12.56,
		},
	}

	for _, test := range tests {
		result, err := calculateArea(test.shape)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != test.want {
			t.Errorf("for %s, expected %.2f, got %.2f", test.name, test.want, result)
		}
	}
}
