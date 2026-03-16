package main

import "fmt"

func main() {
	//firstTask()

	//secondTask()

	thirdTask()
}
func firstTask() {
	type Car struct {
		model       string
		engine      float64
		releaseDate int
	}

	bmw := Car{"m3", 2.0, 2015}
	haval := Car{"f7x", 1.5, 2021}
	mercedes := Car{"gle", 2.0, 2019}

	fmt.Println(bmw)
	fmt.Println(haval)

	fmt.Println("haval произвел замену двигателя")
	haval.engine = 2.0
	fmt.Println(haval)

	fmt.Println(mercedes)
}

type manufacturer int

const (
	apple manufacturer = iota
	samsung
	xiaomi
	honor
	redmi
	huawei
)

type Phone struct {
	manufacturer    manufacturer
	model           string
	memoryGB        int
	backCameraMP    float64
	frontalCameraMP float64
	isGlobalVersion bool
}

func createNewPhone(manufacturer manufacturer, model string, memoryGb int, backCameraMP float64, frontalCameraMP float64, isGlobalVersion bool) Phone {
	if model == "" {
		fmt.Println("Введите модель телефона и попробуйте и снова")
		return Phone{}
	}

	if memoryGb < 64 || memoryGb > 1024 {
		fmt.Println("Объем памяти не может быть меньше 64 гигабайт и больше 1024 гигабайта, попробуйте снова")
		return Phone{}
	}

	if backCameraMP < 2.0 && frontalCameraMP < 2.0 {
		fmt.Println("Телефонов с камерами меньше 2.0 мегапикселей  не выпускают, попробуйте снова")
		return Phone{}
	}

	return Phone{
		manufacturer,
		model,
		memoryGb,
		backCameraMP,
		frontalCameraMP,
		isGlobalVersion,
	}
}

func (p *Phone) upgradePhoneMemory(newValue int) {
	if newValue > 64 && newValue < 1024 {
		p.memoryGB = newValue
	} else {
		fmt.Println("Введите валидное значение от 1 до 1024")
	}
}

func (p *Phone) changeModelName(newName string) {
	if newName != "" {
		p.model = newName
	} else {
		fmt.Println("Имя не может быть пустым")
	}
}

func (p *Phone) setIsGlobalVersion() {
	p.isGlobalVersion = true
}

func (p *Phone) setIsLocalVersion() {
	p.isGlobalVersion = false
}

func secondTask() {
	iPhone16 := createNewPhone(apple, "Iphone 16", 256, 16.0, 8.0, true)
	fmt.Println(iPhone16)

	iPhone16.changeModelName("")
	fmt.Println(iPhone16)
}

type Float struct {
	address string
	square  *float64
	price   int
}

func (f *Float) setFloatPrice(newPrice int) {
	if newPrice > 0 {
		f.price = newPrice
	}
}

func (f *Float) setFloatAddress(newAddress string) {
	if newAddress == "" {
		return
	}

	if f.address != "" && newAddress != f.address {
		fmt.Println("Адрес квартиры изменять нельзя")
	}

	f.address = newAddress
}

func (f *Float) setFloatSquare(newSquare float64) {

	if newSquare < 10.0 {
		fmt.Println("Размер площади квартиры не может быть меньше 10 квадратных метров")
		return
	}

	if f.square != nil {
		fmt.Println("Площадь уже задана и не может быть изменена")
		return
	}

	f.square = &newSquare
}

func createFloat(address string, square float64, price int) Float {
	return Float{
		address: address,
		square:  &square,
		price:   price,
	}
}

func thirdTask() {
	newFloat := createFloat("Улица Пушкина, дом Колотушкина", 35.5, 50000)
	fmt.Println(newFloat)
}
