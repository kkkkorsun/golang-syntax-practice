package models

import "fmt"

type Manufacturer int

const (
	Apple Manufacturer = iota
	Samsung
	Xiaomi
	Honor
	Redmi
	Huawei
)

type Phone struct {
	manufacturer    Manufacturer
	model           string
	memoryGB        int
	backCameraMP    float64
	frontalCameraMP float64
	isGlobalVersion bool
}

func CreateNewPhone(manufacturer Manufacturer, model string, memoryGb int, backCameraMP float64, frontalCameraMP float64, isGlobalVersion bool) Phone {
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

func (p *Phone) ChangeModelName(newName string) {
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
