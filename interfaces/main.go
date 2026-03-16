package main

import (
	"fmt"
	geometry "interfaces/geometry-calculator"
	"interfaces/geometry-calculator/figures"
	"interfaces/notifications"
	"interfaces/notifications/channels"
	"interfaces/payments"
	"interfaces/payments/methods"
)

func main() {
	//theoryFromVideo()

	//firstTask()

	secondTask()
}

func theoryFromVideo() {
	method := methods.NewBankMethod()
	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("Игра", 5)
	paymentModule.Pay("Бургер", 10)
	compId := paymentModule.Pay("Компьютер", 2500)

	paymentModule.Cancel(compId)

	fmt.Println(paymentModule.AllInfo())

}

func firstTask() {
	figure := figures.CreateTriangle(3, 4, 5)
	geometryModule := geometry.NewGeometryModule(figure)

	area := geometryModule.CalculateArea()
	perimeter := geometryModule.CalculatePerimeter()

	fmt.Println("Вычислили площать: ", area)
	fmt.Println("Вычислили периметр: ", perimeter)

}

func secondTask() {
	m := map[notifications.Channel]notifications.NotificationMethod{

		notifications.ChannelEmail: channels.CreateEmailChannel(),
		notifications.ChannelPush:  channels.CreatePushChannel(),
		notifications.ChannelSMS:   channels.CreateSmsChannel(),
	}
	var n []notifications.Notification

	n = append(n,
		notifications.CreateNotification("Какич", notifications.ChannelSMS),
		notifications.CreateNotification("Пукич", notifications.ChannelEmail),
		notifications.CreateNotification("Когда в пубг?", notifications.ChannelPush))

	module := notifications.CreateNotificationModule(m)

	module.Send(n)

}
