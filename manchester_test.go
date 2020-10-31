package manchester

import (
	"log"
	l "log"
	"testing"
	"time"
)

const TestClock = time.Millisecond * 50

const enableTicker = true

type pin struct {
	v bool
}

func (p pin) Get() bool {
	return p.v
}

func TestDecoder(t *testing.T) {

	d := Decoder{}

	var first = true
	var last bool
	count := 0

	sendBool := func(data bool) {
		count++
		if last != data || first {
			d.Decode(data, float32(TestClock/time.Millisecond/2)*float32(count))
			count = 0
		}
		last = data
		first = false
	}
	sendBoolManchester := func(data bool) {
		sendBool(!data)
		sendBool(data)
	}

	sendBytes := func(data []byte) {
		for _, b := range data {
			l.Printf("sending byte: %08b", b)
			for j := 0; j < 8; j++ {
				sendBoolManchester(b>>uint(7-j)&0x01 == 1)
			}
		}
	}

	sendMessage := func(data []byte) {
		sendBytes(data)
	}

	sendBytes([]byte{0b11111111})
	sendBytes([]byte{0b11110100})
	sendMessage([]byte("elliot was here"))

	var decoded []byte
	for {
		if d.Available() == 0 {
			break
		}

		b, _ := d.Read()
		decoded = append(decoded, b)
	}

	log.Printf("decoded message: %s", decoded)
	if string(decoded) != "elliot was here" {
		panic("wrong message decoded")
	}
}
