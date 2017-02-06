# go-roku-ecs
Go library for the Roku External Control Service

External Control Service API
============================
Roku has an amazingly simple and open API for controlling Roku devices over the local network. This library implments a subset of the features decsribed in Roku's [External Control Guide](https://sdkdocs.roku.com/display/sdkdoc/External+Control+Guide).

Usage
=====

Call `roku.Connect()` with the IP address of the Roku device you'd like control, such as `192.168.1.50`. The current IP address assigned to your Roku device can be found by selcting Settings -> Network -> Info. Then call various functions to control the Roku box as described in [roku.go](roku.go).

````go
	package main

	import (
		"github.com/mgoff/go-samsung-exlink"
		"log"
		"time"
	)

	func main() {

		// open the connection to the EX-Link device
		device, err := exlink.Open("/dev/tty.KeySerial1")
		if err != nil { log.Fatal(err) }

		// close the connection at the end
		defer device.Close()

		// turn on the EX-Link device
		err = device.PowerOn()
		if err != nil { log.Fatal(err) }
		time.Sleep(10 * time.Second)

		// switch to HDMI 1
		err = device.SourceHDMI1()
		if err != nil { log.Fatal(err) }
		time.Sleep(10 * time.Second)

		// switch to HDMI 2
		err = device.SourceHDMI2()
		if err != nil { log.Fatal(err) }
		time.Sleep(10 * time.Second)

		// turn off the EX-Link device
		err = device.PowerOff()
		if err != nil { log.Fatal(err) }
	}
````
