package adps9960

// Constants/addresses used for the ADPS9960

// The I2C address which this device listens to.
const Address = 0x39

const (

	/* Gesture parameters */
	GESTURE_THRESHOLD_OUT = 10
	GESTURE_SENSITIVITY_1 = 50
	GESTURE_SENSITIVITY_2 = 20

	/* Error code for returned values */
	ERROR = 0xFF

	/* Acceptable device IDs */
	ID_1 = 0xAB
	ID_2 = 0x9C
	ID_3 = 0xA8

	/* Misc parameters */
	FIFO_PAUSE_TIME = 30 // Wait period (ms) between FIFO reads

	/* APDS-9960 register addresses */
	ENABLE     = 0x80
	ATIME      = 0x81
	WTIME      = 0x83
	AILTL      = 0x84
	AILTH      = 0x85
	AIHTL      = 0x86
	AIHTH      = 0x87
	PILT       = 0x89
	PIHT       = 0x8B
	PERS       = 0x8C
	CONFIG1    = 0x8D
	PPULSE     = 0x8E
	CONTROL    = 0x8F
	CONFIG2    = 0x90
	ID         = 0x92
	STATUS     = 0x93
	CDATAL     = 0x94
	CDATAH     = 0x95
	RDATAL     = 0x96
	RDATAH     = 0x97
	GDATAL     = 0x98
	GDATAH     = 0x99
	BDATAL     = 0x9A
	BDATAH     = 0x9B
	PDATA      = 0x9C
	POFFSET_UR = 0x9D
	POFFSET_DL = 0x9E
	CONFIG3    = 0x9F
	GPENTH     = 0xA0
	GEXTH      = 0xA1
	GCONF1     = 0xA2
	GCONF2     = 0xA3
	GOFFSET_U  = 0xA4
	GOFFSET_D  = 0xA5
	GOFFSET_L  = 0xA7
	GOFFSET_R  = 0xA9
	GPULSE     = 0xA6
	GCONF3     = 0xAA
	GCONF4     = 0xAB
	GFLVL      = 0xAE
	GSTATUS    = 0xAF
	IFORCE     = 0xE4
	PICLEAR    = 0xE5
	CICLEAR    = 0xE6
	AICLEAR    = 0xE7
	GFIFO_U    = 0xFC
	GFIFO_D    = 0xFD
	GFIFO_L    = 0xFE
	GFIFO_R    = 0xFF

	/* Bit fields */
	PON           = 0b00000001
	AEN           = 0b00000010
	PEN           = 0b00000100
	WEN           = 0b00001000
	APSD9960_AIEN = 0b00010000
	PIEN          = 0b00100000
	GEN           = 0b01000000
	GVALID        = 0b00000001

	/* On/Off definitions */
	OFF = 0
	ON  = 1

	/* Acceptable parameters for setMode */
	POWER             = 0
	AMBIENT_LIGHT     = 1
	PROXIMITY         = 2
	WAIT              = 3
	AMBIENT_LIGHT_INT = 4
	PROXIMITY_INT     = 5
	GESTURE           = 6
	ALL               = 7

	/* LED Drive values */
	LED_DRIVE_100MA  = 0
	LED_DRIVE_50MA   = 1
	LED_DRIVE_25MA   = 2
	LED_DRIVE_12_5MA = 3

	/* Proximity Gain (PGAIN) values */
	PGAIN_1X = 0
	PGAIN_2X = 1
	PGAIN_4X = 2
	PGAIN_8X = 3

	/* ALS Gain (AGAIN) values */
	AGAIN_1X  = 0
	AGAIN_4X  = 1
	AGAIN_16X = 2
	AGAIN_64X = 3

	/* Gesture Gain (GGAIN) values */
	GGAIN_1X = 0
	GGAIN_2X = 1
	GGAIN_4X = 2
	GGAIN_8X = 3

	/* LED Boost values */
	LED_BOOST_100 = 0
	LED_BOOST_150 = 1
	LED_BOOST_200 = 2
	LED_BOOST_300 = 3

	/* Gesture wait time values */
	GWTIME_0MS    = 0
	GWTIME_2_8MS  = 1
	GWTIME_5_6MS  = 2
	GWTIME_8_4MS  = 3
	GWTIME_14_0MS = 4
	GWTIME_22_4MS = 5
	GWTIME_30_8MS = 6
	GWTIME_39_2MS = 7

	/* Default values */
	DEFAULT_ATIME          = 219  // 103ms
	DEFAULT_WTIME          = 246  // 27ms
	DEFAULT_PROX_PPULSE    = 0x87 // 16us, 8 pulses
	DEFAULT_GESTURE_PPULSE = 0x89 // 16us, 10 pulses
	DEFAULT_POFFSET_UR     = 0    // 0 offset
	DEFAULT_POFFSET_DL     = 0    // 0 offset
	DEFAULT_CONFIG1        = 0x60 // No 12x wait (WTIME) factor
	DEFAULT_LDRIVE         = LED_DRIVE_100MA
	DEFAULT_PGAIN          = PGAIN_4X
	DEFAULT_AGAIN          = AGAIN_4X
	DEFAULT_PILT           = 0      // Low proximity threshold
	DEFAULT_PIHT           = 50     // High proximity threshold
	DEFAULT_AILT           = 0xFFFF // Force interrupt for calibration
	DEFAULT_AIHT           = 0
	DEFAULT_PERS           = 0x11 // 2 consecutive prox or ALS for int.
	DEFAULT_CONFIG2        = 0x01 // No saturation interrupts or LED boost
	DEFAULT_CONFIG3        = 0    // Enable all photodiodes, no SAI
	DEFAULT_GPENTH         = 40   // Threshold for entering gesture mode
	DEFAULT_GEXTH          = 30   // Threshold for exiting gesture mode
	DEFAULT_GCONF1         = 0x40 // 4 gesture events for int., 1 for exit
	DEFAULT_GGAIN          = GGAIN_4X
	DEFAULT_GLDRIVE        = LED_DRIVE_100MA
	DEFAULT_GWTIME         = GWTIME_2_8MS
	DEFAULT_GOFFSET        = 0    // No offset scaling for gesture mode
	DEFAULT_GPULSE         = 0xC9 // 32us, 10 pulses
	DEFAULT_GCONF3         = 0    // All photodiodes active during gesture
	DEFAULT_GIEN           = 0    // Disable gesture interrupts
)
