package channels

import "fmt"

type Sms struct {
}

func CreateSmsChannel() Sms {
	return Sms{}
}

func (s Sms) Send(text string) bool {
	if len(text) != 0 {
		fmt.Println("Отправлено SMS уведомление с текстом: ", text)
		return true
	}
	return false
}
