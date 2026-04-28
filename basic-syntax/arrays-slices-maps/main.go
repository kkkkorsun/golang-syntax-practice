package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//firstTask()
	//secondTask()
	//thirdTask()
	//fourTask()
	fiveTask()
}

func firstTask() {
	arrayOfStrings := [5]string{"a", "b", "c", "d", "e"}
	arrayOfInt := [3]int{1, 2, 3}

	for i := 0; i <= len(arrayOfStrings); i++ {
		fmt.Println(i)
	}

	for _, v := range arrayOfStrings {
		fmt.Println(v)

	}

	for _, v := range arrayOfInt {
		if v%2 == 0 {
			v += 10
		}
		fmt.Println(v)
	}
}

func secondTask() {
	var helloStrings = []string{"Hello", "world"} // Создание слайса на месте
	var emptySliceOfInt []int                     //Создание слайса без выделения памяти
	sliceOfInt := make([]int, 0)                  //Создание слайса без выделения памяти через make
	sliceOfFloat := make([]float64, 5, 10)        //Создание слайса со значениями по умолчанию

	fmt.Println(helloStrings)
	fmt.Println(emptySliceOfInt)
	fmt.Println(sliceOfInt)
	fmt.Println(sliceOfFloat)

	var firstIndexValue = helloStrings[0]
	fmt.Println("Выводим первый индекс слайса строк: ", firstIndexValue)

	emptySliceOfInt = append(emptySliceOfInt, 10)
	fmt.Println("Добавил значение в слайс с целыми значениями: ", emptySliceOfInt)

	for i := 0; i < 10; i++ {
		var randomNumber = rand.Int()
		fmt.Println("Добавляем значение в массив sliceOfInt: ", randomNumber)

		sliceOfInt = append(sliceOfInt, randomNumber)
		fmt.Println("Текущий размер массива: ", len(sliceOfInt))
		fmt.Println("Текущая емкость массива: ", cap(sliceOfInt))
	}

	fmt.Println("Значения слайса: ", sliceOfInt)

	sliceOfInt[4] = 10

	fmt.Println("Поменял 4 значение слайса на 10, проверяем: ", sliceOfInt)

	test(sliceOfInt)

	fmt.Println("\n")
	fmt.Println(sliceOfInt)
}

// Изменения применились потому что слайсы передают указатель на массив
func test(v []int) {
	v[4] = 14
	fmt.Println(v)
}

func thirdTask() {
	var weatherByDayMap = map[int]int{}

	for i := 1; i <= 31; i++ {
		weatherByDayMap[i] = GenerateRandomNumber(10, 25)
	}
	fmt.Println(weatherByDayMap)
	weatherByDayMap[5] = 20
	fmt.Println("Погода 5 числа изменилась, теперь там: ", weatherByDayMap[5])

	var weatherForYear = map[string]map[int]int{
		"Январь":   {},
		"Февраль":  {},
		"Март":     {},
		"Апрель":   {},
		"Май":      {},
		"Июнь":     {},
		"Июль":     {},
		"Август":   {},
		"Сентябрь": {},
		"Октябрь":  {},
		"Ноябрь":   {},
		"Декабрь":  {},
	}

	for k, v := range weatherForYear {
		if k == "Февраль" {
			for i := 1; i <= 29; i++ {
				v[i] = GenerateRandomNumber(-10, 10)
			}
		} else {
			for i := 1; i < 31; i++ {
				v[i] = GenerateRandomNumber(0, 15)
			}
		}

		fmt.Println("Прогноз погоды: ", k, v)
	}

	//delete(weatherForYear, "Ноябрь")
	//Проверка на значение по умолчанию
	month, ok := weatherForYear["Ноябрь"]
	fmt.Println(month, ok)

}

func GenerateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min // +1 for inclusive max
}

type Dog struct {
	Name     string
	Rating   int
	IsPolite bool
}

func fourTask() {
	competitionParticipants := make([]Dog, 0)

	competitionParticipants = append(competitionParticipants,
		Dog{
			Name:     "Тузик",
			Rating:   0,
			IsPolite: false,
		},
		Dog{
			Name:     "Марсик",
			Rating:   0,
			IsPolite: true,
		},
		Dog{
			Name:     "Жорик",
			Rating:   0,
			IsPolite: true,
		})

	for _, v := range competitionParticipants {
		fmt.Println("Приветствуем участника конкурса: ", v.Name)
	}

	fmt.Println("Все пёс пёсычи прошли выставку, подсчитываем оценки...")

	for i, v := range competitionParticipants {
		if v.IsPolite {
			competitionParticipants[i].Rating++
		}

		fmt.Println("Смотрим на итоговые оценки: ", competitionParticipants[i])
	}
}

type parkingSpace struct {
	ParkingNumber   string
	ParkingCost     int
	ParkingDiscount float64
}

func fiveTask() {
	parking := []parkingSpace{
		{
			ParkingNumber:   "A4",
			ParkingCost:     150,
			ParkingDiscount: 0,
		},
		{
			ParkingNumber:   "B8",
			ParkingCost:     380,
			ParkingDiscount: 0,
		},
		{
			ParkingNumber:   "D32",
			ParkingCost:     900,
			ParkingDiscount: 0,
		},
	}
	fmt.Println("Наша парковка: ", parking)

	for i, _ := range parking {
		if parking[i].ParkingCost < 500 {
			fmt.Println("Парковка дешевле 500: ", parking[i])
		}

		if parking[i].ParkingCost >= 900 {
			parking[i].ParkingDiscount = 10
			fmt.Println("Парковочное место получило скидку", parking[i])
		}

		if parking[i].ParkingDiscount > 0 {
			parking[i].ParkingCost = int(float64(parking[i].ParkingCost) * (1 - float64(parking[i].ParkingDiscount)/100))
		}
	}

	fmt.Println("Наша парковка после действий над ней: ", parking)

}
