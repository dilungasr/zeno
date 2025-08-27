package zeno

import "time"

var apiConfigOptions Options

func ConfigAPI(options Options) {
	if options.Timeout == 0 {
		options.Timeout = time.Second * 50
	}
	// set API access data
	apiConfigOptions = options
}
