package main

import "fmt"

func main() {

	firstCase()

	secondCase()

	thirdCase()

	fourthCase()

}

func firstCase() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("В первом кейсе с делением была получена паника: ", p)
		}
	}()
	a := 10
	b := 0
	c := a / b

	fmt.Println("Результат деления: ", c)
}

func secondCase() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Во втором кейсе с nil была получена паника: ", p)
		}
	}()

	var m map[int]int
	m[0] = 10
}

func thirdCase() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("В третьем кейсе с N+1 при длине N была получена паника: ", p)
		}
	}()
	var s = []int{1, 2, 3, 5}
	fmt.Println("Попытка вывести значение за пределами", s[10])
}

func fourthCase() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("В четвертом кейсе отловили ручную панику: ", p)
		}
	}()

	panic("Тестовая паника")
}
