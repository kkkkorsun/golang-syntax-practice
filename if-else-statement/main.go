package main

import "fmt"

func main() {
	//firstTask()
	//secondTask()
	//thirdTask()
	//fourthTask()
	//fifthTask()
	sixthTask()
}

func firstTask() {
	age := 39
	if age > 30 {
		fmt.Print("Тебе уже больше тридцатки")
	} else {
		fmt.Print("Ты еще молодой!")
	}
}

func secondTask() {
	age := 5

	if age >= 18 {
		fmt.Print("Я продам тебе пиво")
	} else if age <= 12 {
		fmt.Print("Ты чего малой берега совсем попутал пиво покупать?")
	} else {
		fmt.Print("Алкоголь только от 18 лет!")
	}
}

func thirdTask() {
	litres := 2.0

	if litres < 1.0 || litres > 3 {
		fmt.Print("Ты какой-то странный....")
	} else {
		fmt.Print("А ты знаешь золотую середину")
	}
}

func fourthTask() {
	beerDelicious := false
	croutonsDelicious := true

	if beerDelicious && croutonsDelicious {
		fmt.Print("Мы идем гулять")
	} else {
		fmt.Print("Мы не идем гулять")
	}
}

func fifthTask() {
	weather := 3
	if weather != 3 {
		fmt.Print("Мы участвуем в гонке")
	} else {
		fmt.Print("Мы не участвуем в гонке")
	}
}

func sixthTask() {
	isPremiumExist := false
	if !isPremiumExist {
		fmt.Print("Бро купи себе премиум")
	}
}
