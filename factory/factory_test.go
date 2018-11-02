package factory

import "testing"

func TestChannelFactory_CreateChannel(t *testing.T) {

	channelFactory := &ChannelFactory{}

	channel := channelFactory.CreateChannel(1)

	channel.SendMessage()

	channel = channelFactory.CreateChannel(2)

	channel.SendMessage()
}
