package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//firstTask()

	//secondTask()

	//thirdTask()

	//fourthTask()

	fifthTask()
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

func thirdTask() {
	go func() {
		fmt.Println("Вывожу текст из первой анонимной горутины")
	}()

	go func() {
		fmt.Println("Вывожу текст из второй анонимной горутины")
	}()

	go func() {
		fmt.Println("Вывожу текст из третьей анонимной горутины")
	}()

	time.Sleep(1 * time.Second)
}

func fourthTask() {
	chInt := make(chan int)
	chString := make(chan string)
	chFloat := make(chan float64)

	go func() {
		for {
			chInt <- rand.Int()
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for {
			chString <- "Пишу число: " + strconv.Itoa(rand.Int())
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			chFloat <- rand.Float64()
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		select {
		case intValue := <-chInt:
			fmt.Println("Выводим из канала chInt: ", intValue)
		case stringValue := <-chString:
			fmt.Println("Выводим из канала chString: ", stringValue)
		case floatValue := <-chFloat:
			fmt.Println("Выводим из канала chFloat: ", floatValue)

		}
	}

}

func fifthTask() {
	firstCh := make(chan int)
	secondCh := make(chan int)
	delays := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}

	go func(channel chan int) {
		randomIndex := rand.Intn(len(delays))
		fmt.Println("Текущая задержка в первом канале: ", delays[randomIndex])
		time.Sleep(time.Duration(delays[randomIndex]) * time.Millisecond)

		channel <- rand.Int()
	}(firstCh)

	go func(channel chan int) {
		randomIndex := rand.Intn(len(delays))
		fmt.Println("Текущая задержка во втором канале: ", delays[randomIndex])
		time.Sleep(time.Duration(delays[randomIndex]) * time.Millisecond)

		channel <- rand.Int()
	}(secondCh)

	time.Sleep(500 * time.Millisecond)

	select {
	case firstChValue := <-firstCh:
		fmt.Println("Читаем из первого канала", firstChValue)
	case secondChValue := <-secondCh:
		fmt.Println("Читаем из второго канала", secondChValue)
	default:
		fmt.Println("К моменту чтения ничего не было записано в канал")
	}
}
