package main

import (
	"fmt"
	"log"
	"time"

	"machine"

	manchester "github.com/elliots/tinygo-manchester"
	"github.com/elliots/tinygo-manchester/examples/clue/adps9960"
)

func main() {

	time.Sleep(time.Second * 5)
	log.Printf("starting\r\n")

	machine.I2C0.Configure(machine.I2CConfig{})
	d := adps9960.New(machine.I2C0)
	d.Configure()
	adps9960.Debug = false

	log.Printf("sensor configured\r\n")

	m := &manchester.Decoder{}
	manchester.Debug = true

	var raw uint16
	var err error
	var val, lastVal bool
	var now, lastNow time.Time
	var diff float32

	lastNow = time.Now()
	for {
		raw, err = d.ReadAmbientLight()
		if err != nil {
			println("main: light read error: " + err.Error())
		} else {
			val = raw != 0
			if val != lastVal {
				now = time.Now()
				diff = float32(now.Sub(lastNow) / time.Millisecond)
				fmt.Printf("edge after: %fms\r\n", diff)

				m.Decode(val, diff)

				lastVal = val
				lastNow = now
			}
		}
	}

}
