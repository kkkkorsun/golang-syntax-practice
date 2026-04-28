package channels

import "fmt"

type Email struct {
}

func CreateEmailChannel() Email {
	return Email{}
}

func (e Email) Send(text string) bool {

	if len(text) != 0 {
		fmt.Println("Отправлено уведомление по Email с текстом: ", text)
		return true
	}

	return false
}
