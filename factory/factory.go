package factory

import "fmt"

type Channel interface {
	SendMessage()
}

type SmsChannel struct {
}

func (SmsChannel) SendMessage() {
	fmt.Println("sms send message")
}

type EmailChannel struct {
}

func (EmailChannel) SendMessage() {
	fmt.Println("email send message")
}

//工厂模式
type ChannelFactory struct {
}

func (*ChannelFactory) CreateChannel(messageType int) Channel {
	switch messageType {
	case 1:
		return SmsChannel{}
	case 2:
		return EmailChannel{}
	}
	return nil
}
