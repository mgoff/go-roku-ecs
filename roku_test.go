package roku

import (
	"testing"
	"time"
)

// udpate devicePath with your serial port path before running 'go test'
const (
	address = "192.168.1.50"
)

// sleep in between test commands so you can watch the results on your TV
func TestRokuEcp(t *testing.T) {

	// open the connection to the Roku device
	device, err := Connect(address)
	if err != nil {
		t.Error(err)
	}

	// launch Netflix
	if device.LaunchNetflix() != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)

	// launch live TV on CNNGo
	if device.ChannelCNNLive() != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)

	// launch channel CCTV4 on Damai
	if device.ChannelCCTV4() != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)

	// go Home
	if device.KeypressHome() != nil {
		t.Error(err)
	}
}
