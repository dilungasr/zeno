package zeno

import (
	"fmt"
	"log"
)

// structures logging message for zeno API
func zLog(format string, a ...any) (msg string) {
	format = "[zeno-pay] " + format
	msg = fmt.Sprintf(format, a...)
	log.Println(msg)

	return msg
}
