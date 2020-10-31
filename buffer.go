package manchester

const bufferSize = 128

// RingBuffer is ring buffer implementation inspired by post at
// https://www.embeddedrelated.com/showthread/comp.arch.embedded/77084-1.php
// Copied from https://github.com/tinygo-org/tinygo/blob/release/src/machine/buffer.go

type RingBuffer struct {
	rxbuffer [bufferSize]byte
	head     byte
	tail     byte
}

// NewRingBuffer returns a new ring buffer.
func NewRingBuffer() *RingBuffer {
	return &RingBuffer{}
}

// Used returns how many bytes in buffer have been used.
func (rb *RingBuffer) Used() uint8 {
	return uint8(rb.head - rb.tail)
}

// Put stores a byte in the buffer. If the buffer is already
// full, the method will return false.
func (rb *RingBuffer) Put(val byte) bool {
	if rb.Used() != bufferSize {
		rb.head++
		rb.rxbuffer[rb.head%bufferSize] = val
		return true
	}
	return false
}

// Get returns a byte from the buffer. If the buffer is empty,
// the method will return a false as the second value.
func (rb *RingBuffer) Get() (byte, bool) {
	if rb.Used() != 0 {
		rb.tail++
		return rb.rxbuffer[rb.tail%bufferSize], true
	}
	return 0, false
}

// Clear resets the head and tail pointer to zero.
func (rb *RingBuffer) Clear() {
	rb.head = 0
	rb.tail = 0
}
