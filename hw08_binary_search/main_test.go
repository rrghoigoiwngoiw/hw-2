package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		search int
		want   int
	}{
		{
			name:   "standartMid",
			input:  []int{1, 3, 5, 7, 9},
			search: 5,
			want:   2,
		},
		{
			name:   "noElements",
			input:  []int{1, 3, 5, 7, 9},
			search: 10,
			want:   -1,
		},
		{
			name:   "массив из 1 знака",
			input:  []int{1},
			search: 1,
			want:   0,
		},
		{
			name:   "элемент в начале",
			input:  []int{1, 3, 5, 7, 9},
			search: 1,
			want:   0,
		},
		{
			name:   "большой слайс",
			input:  []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39},
			search: 32,
			want:   -1,
		},
		{
			name:   "еще один большой слайс",
			input:  []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39},
			search: 33,
			want:   16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.input, tt.search)
			if got != tt.want {
				t.Errorf("BinarySearch(%v, %d) = %d, want %d", tt.input, tt.search, got, tt.want)
			}
		})
	}
}
