package serialize

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/yaml.v3"
)

type Book struct {
	ID     int     `json:"id" xml:"id" yaml:"id" bson:"id"`
	Title  string  `json:"title" xml:"title" yaml:"title" bson:"title"`
	Author string  `json:"author" xml:"author" yaml:"author" bson:"author"`
	Year   string  `json:"year" xml:"year" yaml:"year" bson:"year"`
	Size   int     `json:"size" xml:"size" yaml:"size" bson:"size"`
	Rate   float64 `json:"rate" xml:"rate" yaml:"rate" bson:"rate"`
	Sample []byte  `json:"sample" xml:"sample" yaml:"sample,flow" bson:"sample"`
}

func JSONSerialize(book Book) (string, error) {
	jsonbook, err := json.Marshal(book)
	if err != nil {
		return "", err
	}
	return string(jsonbook), nil
}

func JSONDeserialize(jsonString string) (Book, error) {
	var jsonbook Book
	err := json.Unmarshal([]byte(jsonString), &jsonbook)
	if err != nil {
		return Book{}, err
	}
	return jsonbook, nil
}

func XMLSerialize(book Book) (string, error) {
	xmlbook, err := xml.MarshalIndent(book, "", " ")
	if err != nil {
		return "", err
	}
	return string(xmlbook), nil
}

func XMLDeserialize(xmlstring string) (Book, error) {
	var xmlbook Book
	err := xml.Unmarshal([]byte(xmlstring), &xmlbook)
	if err != nil {
		return Book{}, err
	}
	return xmlbook, nil
}

func YAMLSerialize(book Book) (string, error) {
	yamlbook, err := yaml.Marshal(book)
	if err != nil {
		return "", err
	}
	return string(yamlbook), nil
}

func YAMLDeserialize(yamlString string) (Book, error) {
	var yamlBook Book
	err := yaml.Unmarshal([]byte(yamlString), &yamlBook)
	if err != nil {
		return Book{}, err
	}
	return yamlBook, nil
}

func GOBSerialize(book Book) (string, error) {
	var buf bytes.Buffer
	gobbook := gob.NewEncoder(&buf)
	err := gobbook.Encode(book)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func GOBDeserialize(gobData []byte) (Book, error) {
	var gobbook Book
	buf := bytes.NewBuffer(gobData)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(&gobbook)
	if err != nil {
		return Book{}, err
	}
	return gobbook, nil
}

func BSONSerialize(book Book) ([]byte, error) {
	bsonbook, err := bson.Marshal(book)
	if err != nil {
		return nil, err
	}
	return bsonbook, nil
}

func BSONDeserialize(bsonData []byte) (Book, error) {
	var bsonbook Book
	err := bson.Unmarshal(bsonData, &bsonbook)
	if err != nil {
		return Book{}, err
	}
	return bsonbook, nil
}
