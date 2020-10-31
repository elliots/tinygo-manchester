package adps9960

import "fmt"

// Debug enables the logging. No suprise there.
var Debug = true

func log(msg string, args ...interface{}) {
	if Debug {
		fmt.Printf("adps9960: "+msg+"\r\n", args...)
	}
}
