package exlink

import (
	"net/http"
	"strconv"
	"time"
)

// common Roku app IDs
const (
	CNNGo	= 65978
	Damai	= 113040
	Netflix	= 12
)

// struct used to hold our Roku connection
type Roku struct {
	address string
	client *http.Client
}

// open a connection to the Roku device
func Connect(address string) (Roku, error) {
	return Roku{address:address, client:&http.Client{}}, nil
}

// common http post method
func (r Roku) Post(path string) error {
    req, err := http.NewRequest("POST", "http://" + r.address + ":8060" + path, nil)
    if err != nil { return err }
	_, err = r.client.Do(req)
	return err
}

// return to Home
func (r Roku) KeypressHome() error {
	return r.Post("/keypress/home")
}

// right button
func (r Roku) KeypressRight() error {
	return r.Post("/keypress/right")
}

// down button
func (r Roku) KeypressDown() error {
	return r.Post("/keypress/down")
}

// select button
func (r Roku) KeypressSelect() error {
	return r.Post("/keypress/select")
}

// launch Netflix
func (r Roku) LaunchNetflix() error {
	return r.Post("/launch/" + strconv.Itoa(Netflix))
}

// launch CNNGo
func (r Roku) LaunchCNN() error {
	return r.Post("/launch/" + strconv.Itoa(CNNGo))
}

// launch Damai
func (r Roku) LaunchDamai() error {
	return r.Post("/launch/" + strconv.Itoa(Damai))
}

// switch to channel CNNLive - subscription required
func (r Roku) ChannelCNNLive() error {
	r.KeypressHome()
	time.Sleep(2 * time.Second)
	r.LaunchCNN()
	time.Sleep(8 * time.Second)
	r.KeypressDown()
	r.KeypressSelect()
	r.KeypressDown()
	r.KeypressSelect()
	return nil
}

// switch to channel CCTV4 on Damai - subscription required
func (r Roku) ChannelCCTV4() error {
	r.KeypressHome()
	time.Sleep(2 * time.Second)
	r.LaunchDamai()
	time.Sleep(8 * time.Second)
	r.KeypressDown()
	r.KeypressRight()
	r.KeypressRight()
	r.KeypressSelect()
	return nil
}