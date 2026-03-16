package figures

import (
	"fmt"
	"math"
)

type Triangle struct {
	a int
	b int
	c int
}

func CreateTriangle(a int, b int, c int) Triangle {
	return Triangle{
		a: a,
		b: b,
		c: c,
	}
}

func (t Triangle) CalculateArea() int {
	fmt.Println("Вычисляем площадь треугольника")

	perimeter := t.CalculatePerimeter() / 2
	return int(math.Sqrt(float64(perimeter * (perimeter - t.a) * (perimeter - t.b) * (perimeter - t.c))))
}

func (t Triangle) CalculatePerimeter() int {
	fmt.Println("Вычисляем периметр треугольника")

	return t.a + t.b + t.c
}
