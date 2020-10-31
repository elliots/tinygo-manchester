package adps9960

import (
	"errors"
)

/*
 * @brief Reads the ambient (clear) light level as a 16-bit value
 *
 * @param[out] val value of the light sensor.
 * @return True if operation successful. False otherwise.
 */
func (d *Device) ReadAmbientLight() (val uint16, err error) {
	data1 := []byte{0}
	if err := d.bus.ReadRegister(uint8(d.Address), CDATAL, data1); err != nil {
		return 0, errors.New("failed to read ambient light: " + err.Error())
	}

	data2 := []byte{0}
	if err := d.bus.ReadRegister(uint8(d.Address), CDATAH, data2); err != nil {
		return 0, errors.New("failed to read ambient light: " + err.Error())
	}

	val = uint16(data1[0])
	val = val + (uint16(data2[0]) << 8)

	log("light value %v", val)

	return val, nil
}

/*
 * @brief Turns ambient light interrupts on or off
 *
 * @param[in] enable 1 to enable interrupts, 0 to turn them off
 * @return True if operation successful. False otherwise.
 */
func (d *Device) SetAmbientLightInterruptEnable(enable uint8) error {
	val := []byte{0}

	/* Read value from ENABLE register */
	if err := d.bus.ReadRegister(uint8(d.Address), ENABLE, val); err != nil {
		return errors.New("failed to get ambient light interrupt register: " + err.Error())
	}

	/* Set bits in register to given value */
	enable &= 0b00000001
	enable = enable << 4
	val[0] &= 0b11101111
	val[0] |= enable

	/* Write register value back into ENABLE register */
	if err := d.bus.WriteRegister(uint8(d.Address), ENABLE, val); err != nil {
		return errors.New("failed to set ambient light interrupt register: " + err.Error())
	}

	return nil
}

/*
 * @brief Clears the ambient light interrupt
 *
 * @return True if operation completed successfully. False otherwise.
 */
func (d *Device) ClearAmbientLightInterrupt() error {
	if err := d.bus.ReadRegister(uint8(d.Address), AICLEAR, []byte{0}); err != nil {
		return errors.New("failed to clear ambient light interrupt: " + err.Error())
	}

	return nil
}

// GetInterrupt reads the state of the triggered interrupts
func (d *Device) GetInterrupts() []uint8 {
	data := []uint8{0}
	d.bus.ReadRegister(uint8(d.Address), STATUS, data)
	log("interrupts: %08b", data[0])
	return data
}
