package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"module_packages/models"
)

func main() {
	iPhone16 := models.CreateNewPhone(models.Apple, "Iphone 16", 256, 16.0, 8.0, true)
	fmt.Println(iPhone16)

	iPhone16.ChangeModelName("")
	pp.Println(iPhone16)
}
