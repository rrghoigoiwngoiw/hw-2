package shapes

import "errors"

type Shape interface { // интерфейс
	Area() float64 // метод
}

type Rectangle struct {
	Weight, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Weight * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("нет фигуры")
	}
	return shape.Area(), nil
}
