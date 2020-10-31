package adps9960

import "machine"

type Device struct {
	bus     machine.I2C
	Address uint16
}

// New only creates the Device object, it does not touch the device.
func New(bus machine.I2C) Device {
	return Device{bus, Address}
}

// Configure the device (currently, just enables everything)
func (d *Device) Configure() {
	d.bus.WriteRegister(uint8(d.Address), ENABLE, []byte{ALL})
}
