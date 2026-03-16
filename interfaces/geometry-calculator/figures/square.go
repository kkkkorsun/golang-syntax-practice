package figures

import (
	"fmt"
)

type Square struct {
	sideSize int
}

func CreateNewSquare(sideSize int) Square {
	return Square{
		sideSize: sideSize,
	}
}

func (s Square) CalculateArea() int {
	fmt.Println("Вычисляем площадь квадрата")

	return s.sideSize * s.sideSize
}

func (s Square) CalculatePerimeter() int {
	fmt.Println("Вычисляем периметр квадрата")

	return 4 * s.sideSize
}
