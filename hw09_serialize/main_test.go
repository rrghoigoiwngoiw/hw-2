package serialize

import (
	"bytes"
	"reflect"
	"testing"
)

func TestJSONSerialize(t *testing.T) {
	book := Book{ID: 1, Title: "Test Book", Author: "Author", Year: "2024", Size: 100, Rate: 4.5, Sample: []byte("Sample")}
	expected := `{"id":1,"title":"Test Book","author":"Author","year":"2024","size":100,"rate":4.5,"sample":"U2FtcGxl"}`

	result, err := JSONSerialize(book)
	if err != nil {
		t.Fatalf("JSONSerialize failed: %v", err)
	}

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestJSONDeserialize(t *testing.T) {
	jsonString := `{"id":1,"title":"Test Book","author":"Author","year":"2024","size":100,"rate":4.5,"sample":"U2FtcGxl"}`
	expected := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   "2024",
		Size:   100,
		Rate:   4.5,
		Sample: []byte("Sample"),
	}

	result, err := JSONDeserialize(jsonString)
	if err != nil {
		t.Fatalf("JSONDeserialize failed: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %+v, Got: %+v", expected, result)
	}
}

func TestXMLSerialize(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   "2024",
		Size:   100,
		Rate:   4.5,
		Sample: []byte("Sample"),
	}
	expected := `<Book>
 <id>1</id>
 <title>Test Book</title>
 <author>Author</author>
 <year>2024</year>
 <size>100</size>
 <rate>4.5</rate>
 <sample>Sample</sample>
</Book>`

	result, err := XMLSerialize(book)
	if err != nil {
		t.Fatalf("XMLSerialize failed: %v", err)
	}

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestXMLDeserialize(t *testing.T) {
	xmlString := `<Book>
 <id>1</id>
 <title>Test Book</title>
 <author>Author</author>
 <year>2024</year>
 <size>100</size>
 <rate>4.5</rate>
 <sample>Sample</sample>
</Book>`
	expected := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   "2024",
		Size:   100,
		Rate:   4.5,
		Sample: []byte("Sample"),
	}

	result, err := XMLDeserialize(xmlString)
	if err != nil {
		t.Fatalf("XMLDeserialize failed: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %+v, Got: %+v", expected, result)
	}
}

func TestYAMLSerialize(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   "2024",
		Size:   100,
		Rate:   4.5,
		Sample: []byte("Sample"),
	}
	expected := `id: 1
title: Test Book
author: Author
year: "2024"
size: 100
rate: 4.5
sample: [83, 97, 109, 112, 108, 101]
`

	result, err := YAMLSerialize(book)
	if err != nil {
		t.Fatalf("YAMLSerialize failed: %v", err)
	}

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestYAMLDeserialize(t *testing.T) {
	yamlString := `id: 1
title: Test Book
author: Author
year: "2024"
size: 100
rate: 4.5
`
	expected := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   "2024",
		Size:   100,
		Rate:   4.5,
	}

	result, err := YAMLDeserialize(yamlString)
	if err != nil {
		t.Fatalf("YAMLDeserialize failed: %v", err)
	}

	if !bytes.Equal(result.Sample, expected.Sample) {
		t.Errorf("Field 'Sample' mismatch. Expected: %v, Got: %v", expected.Sample, result.Sample)
	}

	result.Sample = nil
	expected.Sample = nil
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %+v, Got: %+v", expected, result)
	}
}

func TestGOBSerialize(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   "2024",
		Size:   100,
		Rate:   4.5,
		Sample: []byte("Sample"),
	}
	_, err := GOBSerialize(book)
	if err != nil {
		t.Fatalf("GOBSerialize failed: %v", err)
	}
}

func TestGOBDeserialize(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   "2024",
		Size:   100,
		Rate:   4.5,
		Sample: []byte("Sample"),
	}
	gobData, _ := GOBSerialize(book)
	buf := bytes.NewBuffer([]byte(gobData))
	result, err := GOBDeserialize(buf.Bytes())
	if err != nil {
		t.Fatalf("GOBDeserialize failed: %v", err)
	}

	if !reflect.DeepEqual(result, book) {
		t.Errorf("Expected: %+v, Got: %+v", book, result)
	}
}

func TestBSONSerialize(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   "2024",
		Size:   100,
		Rate:   4.5,
		Sample: []byte("Sample"),
	}
	_, err := BSONSerialize(book)
	if err != nil {
		t.Fatalf("BSONSerialize failed: %v", err)
	}
}

func TestBSONDeserialize(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   "2024",
		Size:   100,
		Rate:   4.5,
		Sample: []byte("Sample"),
	}
	bsonData, _ := BSONSerialize(book)
	result, err := BSONDeserialize(bsonData)
	if err != nil {
		t.Fatalf("BSONDeserialize failed: %v", err)
	}

	if !reflect.DeepEqual(result, book) {
		t.Errorf("Expected: %+v, Got: %+v", book, result)
	}
}
