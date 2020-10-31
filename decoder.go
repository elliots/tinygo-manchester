package manchester

import "log"

type Status int

var Debug = false

const (
	PREAMBLE Status = iota //waiting for checksum
	SYNC                   //synced signal
	CHECK                  //passed checksum
	RESYNC                 //lost sync
)

type Decoder struct {
	Status Status

	T float32 //timing of a half bit
	i uint8   //count edges on the preamble stage
	w bool    //keep track of the 2T timing

	bits struct { //used to store bits until a full byte is received
		buffer byte
		pos    uint8
	}

	buffer RingBuffer
}

/*
Decode is called every falling or rising edge. This function performs the decoding and fills the buffer.
bool s -> current logic state
uint16 t -> time in milliseconds since last transition
*/
func (d *Decoder) Decode(s bool, t float32) {

	d.i++

	if d.Status == RESYNC { //if the RESYNC flag is set reset everything and wait for a preamble and resync
		d.i = 0
		d.T = 0
		d.w = false

		d.bits.buffer = 0x0 //clear bits buffer
		d.bits.pos = 0

		d.Status = PREAMBLE

		if Debug {
			log.Printf("<RESYNC>\r\n")
		}
	}

	addBit := func(s bool) {
		if s {
			d.bits.buffer |= 1 << (7 - d.bits.pos)
		}
		d.bits.pos++
	}

	//preamble
	if d.Status == PREAMBLE {

		if Debug {
			log.Printf("preamble %t %f\r\n", s, t)
		}

		if d.i > 3 {

			if d.T == 0 { //if T is undefined, use the first t as T
				d.T = t
			} else if t < d.T*0.75 { //found a much smaller T, use that instead
				d.T = t
			} else if t > d.T*0.75 && t <= d.T*1.5 { //just T
			} else if t > d.T*1.5 && t < d.T*2.5 { //first 2T is a sync signal
				d.Status = SYNC
				d.i = 0

				d.bits.pos = 4 //checksum start at the 4th bit
				//
				addBit(s)

				if Debug {
					log.Printf("SYNC!\r\n")
				}

			} else { //found much longer T
				d.Status = RESYNC
			}
		}

	} else { // SYNCED start to decode to buffer

		if t > d.T*2.5 || t < d.T*0.75 { //timming is very off
			d.Status = RESYNC
		}

		//decode
		if t > d.T*1.5 { //2T
			addBit(s)
		} else { //T
			if !d.w { //first t
				d.w = true
			} else { //second t
				addBit(s)
				d.w = false
			}
		}

		//do a CHECKSUM after sync signal
		if d.Status == SYNC {

			if d.bits.pos > 7 { // wait for a complete byte
				if d.bits.buffer == 0b0100 { //checksum
					d.Status = CHECK

					if Debug {
						log.Printf("CHECK!\r\n")
					}

					//clear
					d.bits.buffer = 0x0
					d.bits.pos = 0
				} else { //invalid checksum
					d.Status = RESYNC

					if Debug {
						log.Printf("<INVALID CHECKSUM> expected: %08b actual: %08b\r\n", 0b0100, d.bits.buffer)
					}
				}
			}
			//start pushing bytes to out buffer
		} else if d.Status == CHECK {
			//received a full byte
			if d.bits.pos > 7 {
				if Debug {
					log.Printf("GOT BYTE: %08b\r\n", d.bits.buffer)
				}
				if !d.buffer.Put(d.bits.buffer) && Debug {
					log.Printf("Warning: manchester decoder buffer is full")
				}

				d.bits.buffer = 0x0 //clear
				d.bits.pos = 0
			}
		}
	}

}

/* Return the number of bytes in the buffer */
func (d *Decoder) Available() uint8 {
	return d.buffer.Used()
}

/* Used to read the buffer */
func (d *Decoder) Read() (byte, bool) {
	return d.buffer.Get()
}
