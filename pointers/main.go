package main

import "fmt"

func main() {
	//firstTask()
	//secondTask()

	//thirdTask()
	fourTask()
}

func firstTask() {
	textToOutput := "Привет!"
	intToOutput := 10
	floatToOutput := 3.14
	boolToOutput := true

	printString(&textToOutput)
	printInt(&intToOutput)
	printFloat(&floatToOutput)
	printBool(&boolToOutput)
}

func printString(pointer *string) {
	fmt.Println("Адрес указателя на строку: ", pointer)
	fmt.Println("Значение строки: ", *pointer)
}

func printInt(pointer *int) {
	fmt.Println("Адрес указателя на число: ", pointer)
	fmt.Println("Значение числа: ", *pointer)
}

func printBool(pointer *bool) {
	fmt.Println("Адрес указателя на булеву переменную: ", pointer)
	fmt.Println("Значение булевой переменной: ", *pointer)
}

func printFloat(pointer *float64) {
	fmt.Println("Адрес указателя на число с плавающей точкой: ", pointer)
	fmt.Println("Значение числа с плавающей точкой: ", *pointer)
}

func secondTask() {
	intToOutput := 10
	printIntWithChange(&intToOutput)
}

func printIntWithChange(pointer *int) {
	fmt.Println("Значение числа до изменения: ", *pointer)
	*pointer = 12
	fmt.Println("Значение числа после изменения: ", *pointer)
}

func thirdTask() {
	//textToOutput := "Привет!"
	var nilTExtToOutput *string

	if nilTExtToOutput != nil {
		fmt.Println("Текст: ", nilTExtToOutput)
	} else {
		fmt.Println("Передан указатель на nil")
	}
}

func fourTask() {
	//boobsToNeed := 3.3
	var boobsToNeed float64
	fmt.Println("Я телочка и хочу грудь: ", boobsToNeed)
	plasticOperation(&boobsToNeed)
}

func plasticOperation(boobsSize *float64) {
	if boobsSize != nil {
		fmt.Println("Пусть сформирует запрос нормально, а то ниче не понятно")
		return
	}
	fmt.Println("Телочка пришла в клинику и ее запрос: ", *boobsSize)
	*boobsSize++
	fmt.Println("Упс, я перепутал импланты и теперь она будет хвастаться размером: ", *boobsSize)
}
