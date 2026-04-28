package main

import "fmt"

func main() {
	//firstTask()
	secondTask()
}

func firstTask() {
	defer func() {
		fmt.Println("Коко")
	}()

	defer func() {
		fmt.Println("Шнейле")
	}()

	fmt.Println("Фа")
}

//Начинаем тусу
//testOne
//deferTestOne
//Пуси джуси на тусе
//testTwo
//deferTestTwo
//Гучи мами на тусе
//testThree
//deferTestThree
//Деньги деньги мами мами на тусе
//В моей жопе имплатны и в моих сиськах импланты они мне платят и платят

func secondTask() {
	defer func() {
		fmt.Println("В моей жопе имплатны и в моих сиськах импланты они мне платят и платят")
	}()
	fmt.Println("Начинаем тусу")

	testOne()

	fmt.Println("Пуси джуси на тусе")

	testTwo()

	fmt.Println("Гучи мами на тусе")

	testThree()

	fmt.Println("Деньги деньги мами мами на тусе")
}

func testOne() {
	defer func() {
		fmt.Println("defer test One")
	}()

	fmt.Println("Test one func")
}

func testTwo() {
	defer func() {
		fmt.Println("defer test two")
	}()

	fmt.Println("Test two func")
}

func testThree() {
	defer func() {
		fmt.Println("defer test three")
	}()

	fmt.Println("Test three func")
}
