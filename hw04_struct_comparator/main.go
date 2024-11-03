package main

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func NewBook(id int, title string, author string, year int, size int, rate float64) *Book {
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

func (b *Book) Rate() float64 {
	return b.rate
}

type CompareMode int

const (
	ByYear CompareMode = iota
	BySize
	ByRate
)

type BookComparator struct {
	mode CompareMode
}

func NewBookComparator(mode CompareMode) *BookComparator {
	return &BookComparator{mode: mode}
}

func (bc BookComparator) Compare(book1, book2 Book) bool {
	switch bc.mode {
	case ByYear:
		return book1.year > book2.year
	case BySize:
		return book1.size > book2.size
	case ByRate:
		return book1.rate > book2.rate
	default:
		return false
	}
}

func main() {
	var id, year, size int
	var title, author string
	var rate float64

	fmt.Println("параметры для первой книги (ID, Title, Author, Year, Size, Rate):")
	fmt.Scanf("%d %s %s %d %d %f", &id, &title, &author, &year, &size, &rate)
	book1 := NewBook(id, title, author, year, size, rate)

	fmt.Println("параметры для второй книги (ID, Title, Author, Year, Size, Rate):")
	fmt.Scanf("%d %s %s %d %d %f", &id, &title, &author, &year, &size, &rate)
	book2 := NewBook(id, title, author, year, size, rate)

	yearComparator := NewBookComparator(ByYear)
	sizeComparator := NewBookComparator(BySize)
	rateComparator := NewBookComparator(ByRate)

	fmt.Println("Compare by Year:", yearComparator.Compare(*book1, *book2))

	fmt.Println("Compare by Size:", sizeComparator.Compare(*book1, *book2))

	fmt.Println("Compare by Rate:", rateComparator.Compare(*book1, *book2))
}
