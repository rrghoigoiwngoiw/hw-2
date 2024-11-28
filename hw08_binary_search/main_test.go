package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		search int
		want   bool
	}{
		{
			name:   "standartMid",
			input:  []int{1, 3, 5, 7, 9},
			search: 5,
			want:   true,
		},
		{
			name:   "noElements",
			input:  []int{1, 3, 5, 7, 9},
			search: 10,
			want:   false,
		},
		{
			name:   "не отсортированн",
			input:  []int{7, 2, 5, 8, 12},
			search: 7,
			want:   true,
		},
		{
			name:   "массив из 1 знака",
			input:  []int{1},
			search: 1,
			want:   true,
		},
		{
			name:   "отрицательные значения",
			input:  []int{-1, -5, 3, -7, 15},
			search: -5,
			want:   true,
		},
		{
			name:   "элемент в начале",
			input:  []int{1, 3, 5, 7, 9},
			search: 1,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.input, tt.search)
			if got != tt.want {
				t.Errorf("BinarySearch(%v, %d) = %t, want %t", tt.input, tt.search, got, tt.want) //+
			}
		})
	}
}
