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
		return
	}

	// launch Netflix
	err = device.LaunchNetflix()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(10 * time.Second)

	// launch live TV on CNNGo
	err = device.ChannelCNNLive()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(10 * time.Second)

	// launch channel CCTV4 on Damai
	err = device.ChannelCCTV4()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(10 * time.Second)

	// go Home
	err = device.KeypressHome()
	if err != nil {
		t.Error(err)
		return
	}
}