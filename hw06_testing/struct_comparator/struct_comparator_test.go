package structs_comparator

import (
	"testing"
)

func TestBookMethods(t *testing.T) {
	// Создаем объект книги
	book := NewBook(1, "Test Title", "Test Author", 2020, 300, 4.5)

	// Проверяем методы установки и получения значений
	book.SetID(2)
	if book.ID() != 2 {
		t.Errorf("expected ID to be 2, got %d", book.ID())
	}

	book.SetTitle("Updated Title")
	if book.Title() != "Updated Title" {
		t.Errorf("expected Title to be 'Updated Title', got %s", book.Title())
	}

	book.SetAuthor("Updated Author")
	if book.Author() != "Updated Author" {
		t.Errorf("expected Author to be 'Updated Author', got %s", book.Author())
	}

	book.SetYear(2021)
	if book.Year() != 2021 {
		t.Errorf("expected Year to be 2021, got %d", book.Year())
	}

	book.SetSize(400)
	if book.Size() != 400 {
		t.Errorf("expected Size to be 400, got %d", book.Size())
	}

	book.SetRate(4.9)
	if book.Rate() != 4.9 {
		t.Errorf("expected Rate to be 4.9, got %f", book.Rate())
	}
}

func TestBookComparator(t *testing.T) {
	book1 := NewBook(1, "Book One", "Author A", 2015, 200, 4.0)
	book2 := NewBook(2, "Book Two", "Author B", 2020, 300, 4.5)

	tests := []struct {
		name       string
		mode       CompareMode
		book1      *Book
		book2      *Book
		wantResult bool
	}{
		{
			name:       "Compare by Year (book2 is newer)",
			mode:       ByYear,
			book1:      book1,
			book2:      book2,
			wantResult: false,
		},
		{
			name:       "Compare by Size (book2 is larger)",
			mode:       BySize,
			book1:      book1,
			book2:      book2,
			wantResult: false,
		},
		{
			name:       "Compare by Rate (book2 has higher rate)",
			mode:       ByRate,
			book1:      book1,
			book2:      book2,
			wantResult: false,
		},
		{
			name:       "Compare by Year (same year)",
			mode:       ByYear,
			book1:      &Book{id: 3, year: 2020},
			book2:      &Book{id: 4, year: 2020},
			wantResult: false,
		},
		{
			name:       "Compare by Rate (same rate)",
			mode:       ByRate,
			book1:      &Book{id: 3, rate: 4.5},
			book2:      &Book{id: 4, rate: 4.5},
			wantResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			comparator := NewBookComparator(test.mode)
			result := comparator.Compare(*test.book1, *test.book2)

			if result != test.wantResult {
				t.Errorf("mode: %v, expected: %v, got: %v", test.mode, test.wantResult, result)
			}
		})
	}
}
