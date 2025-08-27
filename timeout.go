package zeno

import (
	"time"
)

// timeoutStatus checks if the timeout has been reached
func timeoutStatus(orderID string, timeoutFn func(orderID string)) {
	if timeoutFn == nil {
		return
	}

	timeoutTimer := time.NewTimer(apiConfigOptions.Timeout)
	defer timeoutTimer.Stop()

	// wait for timeout
	<-timeoutTimer.C
	zLog("Payment for order %v has timed out", orderID)
	// call timeout function callback
	timeoutFn(orderID)
}
