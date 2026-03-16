package main

import (
	"fmt"
	"math"
)

func main() {
	//firstTask()
	//secondTask()
	thirdTask()
}

func firstTask() {
	outputString("Привет!", 10)
}

func outputString(value string, repeatCount int) {
	for i := 0; i <= repeatCount; i++ {
		fmt.Println(value)
	}
}

func secondTask() {
	getSquareRoot(1, 2, 1)
}

func getSquareRoot(a, b, c int) {
	d := int(math.Pow(float64(b), 2)) - 4*a*c

	if d < 0 {
		fmt.Println("Корней нет")
	} else if d == 0 {
		fmt.Println("Один корень, вычисляем")

		x := -b / (2 * a)
		fmt.Println("Уранение имеет корень: ", x)
	} else {
		fmt.Println("Два корня, вычисляем")

		firstX := (-float64(b) + math.Sqrt(float64(d))) / (2 * float64(a))
		secondX := (-float64(b) - math.Sqrt(float64(d))) / (2 * float64(a))

		fmt.Println("Первый корень: ", firstX)
		fmt.Println("Второй корень: ", secondX)
	}
}

func thirdTask() {
	fmt.Println(polynomialFunc(4))
}

func polynomialFunc(value int) float64 {
	return math.Pow(float64(value), 4) / 2
}
