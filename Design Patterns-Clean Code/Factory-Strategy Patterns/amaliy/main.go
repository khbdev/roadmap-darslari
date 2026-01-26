package main

import "fmt"


type Notifaction interface {
	Send(message string)
}


type EmailNotifaction struct{}

func (e EmailNotifaction) Send(message string) {
	fmt.Println("Email yuborildi", message)
}

type SmsNotifaction struct {}

func (s SmsNotifaction) Send(message string) {
		fmt.Println("Sms yuborildi", message)
}


func NotifactionFactory(notifactiontype string) Notifaction {
	if notifactiontype == "email" {
		return  EmailNotifaction{}
	}
	if notifactiontype == "sms" {
		return  SmsNotifaction{}
	}
	return  nil
}

func main(){
	notifaction := NotifactionFactory("email")
	notifaction.Send("Salom")
}