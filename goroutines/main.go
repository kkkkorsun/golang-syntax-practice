package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//firstTask()

	secondTask()
}

func firstTask() {
	go goroutineTest(1)
	go goroutineTest(2)
	go goroutineTest(3)

	time.Sleep(5000 * time.Millisecond)

}

func goroutineTest(num int) {
	for i := 0; i < 5; i++ {
		fmt.Println("\nЯ горутина № ", num, "делаю вывод на экран: ", i, " раз")
		time.Sleep(time.Second)
	}
}

func secondTask() {
	ch := make(chan int)

	go generateRandomNumber(ch)

	go generateRandomNumber(ch)

	go generateRandomNumber(ch)

	for i := 0; i < 3; i++ {
		fmt.Println("Получили из канала: ", <-ch)
	}
}

func generateRandomNumber(ch chan<- int) {
	num := rand.Int()
	fmt.Println("Метод сгенерировал число: ", num)
	ch <- num
}
