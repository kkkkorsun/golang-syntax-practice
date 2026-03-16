package notifications

import "fmt"

type Channel string

const (
	ChannelEmail Channel = "email"
	ChannelSMS   Channel = "sms"
	ChannelPush  Channel = "push"
)

type NotificationMethod interface {
	Send(text string) bool
}

type Notification struct {
	Text    string
	Channel Channel
}

func CreateNotification(text string, channel Channel) Notification {
	return Notification{
		Text:    text,
		Channel: channel,
	}
}

type NotificationModule struct {
	notificationMethods map[Channel]NotificationMethod
}

func CreateNotificationModule(notificationMethods map[Channel]NotificationMethod) *NotificationModule {

	var tempMap = make(map[Channel]NotificationMethod)

	for k, v := range notificationMethods {
		tempMap[k] = v
	}

	return &NotificationModule{
		notificationMethods: tempMap}
}

func (n *NotificationModule) Send(notifications []Notification) {

	for i := 0; i < len(notifications); i++ {

		channel := notifications[i].Channel
		message := notifications[i].Text

		if len(message) == 0 {
			fmt.Println("Пустое сообщение не может быть отправлено")
			continue
		}

		method, ok := n.notificationMethods[channel]
		if !ok {
			fmt.Println("Данного канала не существует в методах рассылки")
			continue
		}

		if method.Send(message) {
			fmt.Println("Сообщение успешно отправлено")
		} else {
			fmt.Println("Ошибка при отправки сообщения")
		}
	}
}
