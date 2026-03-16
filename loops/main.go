package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//firstTask()
	//secondTask()
	thirdTask()
}

func firstTask() {
	output := "коко шнейле ватафа"
	for i := 0; i < 10; i++ {
		fmt.Printf(output + "\n")
	}
}

func secondTask() {
	var score int
	for i := 0; i <= 10; i++ {
		fmt.Println("Начинаем игру, текущее количество очков: ", score)
		fmt.Println("\n")
		fmt.Println("Залетаем в уровень: ", i)
		fmt.Println("\n")
		score += 5
		fmt.Println("Успешно пролетели уровень")
		fmt.Println("\n")
	}

	fmt.Println("Количество заработанных очков после игры: ", score)
}

func thirdTask() {
	var randLevelCount = GenerateRandomNumber(10, 20)
	var randLoseValue = GenerateRandomNumber(5, 20)

	fmt.Println("Сгенерировано  уровней: ", randLevelCount)
	fmt.Println("Уровень на котором проиграешь: ", randLoseValue)
	var score int
	for i := 0; i <= randLevelCount; i++ {
		if randLoseValue == i {
			fmt.Println("Упс, ты не пролетел уровень, игра окончена")
			break
		}
		fmt.Println("Начинаем игру, текущее количество очков: ", score)
		fmt.Println("\n")
		fmt.Println("Залетаем в уровень: ", i)
		fmt.Println("\n")
		score += 5
		fmt.Println("Успешно пролетели уровень")
		fmt.Println("\n")
	}

	fmt.Println("Количество заработанных очков после игры: ", score)

}
func GenerateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min // +1 for inclusive max
}
