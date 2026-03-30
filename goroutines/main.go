package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//firstTask()

	//secondTask()

	//thirdTask()

	//fourthTask()

	//fifthTask()

	//sixthTask()

	//seventhTask()

	//eighthTask()

	//ninthTask()

	//tenthTask()

	eleventhTask()
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

func sixthTask() {
	//var chInt chan int
	//writeToNilChannel(chInt)
	//readFromNilChannel(chInt)

	nonBufferedChan := make(chan int)

	go func(ch chan int) {
		ch <- 1
		close(ch)
	}(nonBufferedChan)
	//close(nonBufferedChan)

	for {
		if v, ok := <-nonBufferedChan; ok {
			fmt.Println(v)
		}
	}

}

func writeToNilChannel(chInt chan<- int) {
	chInt <- rand.Int()
}

func readFromNilChannel(chInt <-chan int) {
	fmt.Println("Попытка чтения из канала", <-chInt)
}

func seventhTask() {
	peopleOpinions := make(chan string)

	peopleCount := rand.Intn(10)
	fmt.Println("Количество человек:", peopleCount)

	delays := []int{300, 400, 500, 600, 700}
	delay := delays[rand.Intn(len(delays))]
	fmt.Println("Задержка:", delay)

	go func(ch chan<- string) {
		defer close(ch)

		for i := 0; i < peopleCount; i++ {
			time.Sleep(time.Duration(delay) * time.Millisecond)
			ch <- GenerateString(5)
		}
	}(peopleOpinions)

	for opinion := range peopleOpinions {
		fmt.Println(opinion)
	}
}
func GenerateString(length int) string {

	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func eighthTask() {
	parentCxt, parentCancel := context.WithCancel(context.Background())
	middleCxt, middleCancel := context.WithCancel(parentCxt)
	childCxt, childCancel := context.WithCancel(middleCxt)

	go childFunc(childCxt)
	go middleFunc(middleCxt)
	go parentFunc(parentCxt)

	time.Sleep(2 * time.Second)
	parentCancel()

	time.Sleep(3 * time.Second)
	childCancel()

	time.Sleep(1 * time.Second)
	middleCancel()

}

func parentFunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершили parentFunc")
			return
		default:
			fmt.Println("Выполняется parentFunc")
		}

		time.Sleep(1 * time.Second)
	}
}

func middleFunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершили middleFunc")
			return
		default:
			fmt.Println("Выполняется middleFunc")
		}

		time.Sleep(1 * time.Second)
	}
}

func childFunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершили childFunc")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Выполняется childFunc")
		}
	}
}

func ninthTask() {
	delays := []int{500, 600, 700, 800, 900, 1000}

	workersCount := rand.Intn(10) + 1
	fmt.Println("Количество работников: ", workersCount)

	var wg sync.WaitGroup

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go WaterTheGround(i, delays[rand.Intn(len(delays))], &wg)
	}

	wg.Wait()
	fmt.Println("Все рабочие завершили свою работу")

}

func WaterTheGround(id int, delay int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Огородник: ", id, "Ушел выполнять свою работу на: ", delay, "мс")
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func tenthTask() {

	//var voicesInElectionsCount int
	var voicesInElectionsCountAtomic atomic.Int32

	var wg sync.WaitGroup
	//var mtx sync.Mutex

	electorsCount := rand.Intn(10000) + 1
	fmt.Println("Ожидалось голосов: ", electorsCount)

	for i := 0; i < electorsCount; i++ {
		wg.Add(1)

		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done()
			voicesInElectionsCountAtomic.Add(1)

			//mtx.Lock()
			//voicesInElectionsCount++
			//mtx.Unlock()

			//fmt.Println("Избиратель: ", id, " проголосовал")
		}(i, &wg)
	}

	wg.Wait()

	fmt.Println("Получено голосов: ", voicesInElectionsCountAtomic.Load())
}

func eleventhTask() {
	var lettersFromSubscribers []string
	var wg sync.WaitGroup
	var mtx sync.Mutex

	subsCount := rand.Intn(100000) + 1
	fmt.Println("Ожидаемое количество писем", subsCount)

	for i := 0; i < subsCount; i++ {

		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			mtx.Lock()
			lettersFromSubscribers = append(lettersFromSubscribers, GenerateString(5))
			mtx.Unlock()

		}(&wg)
	}

	wg.Wait()

	fmt.Println("Получено писем: ", len(lettersFromSubscribers))

}

func twelfthTask() {
	keystoreMap := map[string]string{}
	
}
