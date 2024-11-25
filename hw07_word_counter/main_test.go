package countwords

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{
			name:  "обычная ситуация",
			input: "Привет, мир!., Это тестовая строка. Привет, мир!",
			want: map[string]int{
				"привет":   2,
				"мир":      2,
				"это":      1,
				"тестовая": 1,
				"строка":   1,
			},
		},
		{
			name:  "пустая строка",
			input: "",
			want:  map[string]int{},
		},
		{
			name:  "разный регистр",
			input: "Hello hello HELLO",
			want: map[string]int{
				"hello": 3,
			},
		},
		{
			name:  "только знаки",
			input: "!!! ??? ,, ;; ::",
			want:  map[string]int{},
		},
		{
			name:  "знаки и слова",
			input: "Go-lang, Go? lang! Go, lang.",
			want: map[string]int{
				"go":   3,
				"lang": 3,
			},
		},
		{
			name:  "цифры",
			input: "123 456 123",
			want: map[string]int{
				"123": 2,
				"456": 1,
			},
		},
		{
			name:  "иероглифы",
			input: "Привет привет привет! 世界 世界",
			want: map[string]int{
				"привет": 3,
				"世界":     2,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountWords(test.input)
			if !reflect.DeepEqual(result, test.want) {
				t.Errorf("For input: %q\nExpected: %v\nGot: %v", test.input, test.want, result)
			}
		})
	}
}
