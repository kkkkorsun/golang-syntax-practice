package channels

import "fmt"

type Push struct {
}

func CreatePushChannel() Push {
	return Push{}
}

func (p Push) Send(text string) bool {
	if len(text) != 0 {
		fmt.Println("Отправлено Push уведомление с текстом: ", text)
		return true
	}
	return false
}
