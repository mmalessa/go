package mmessenger

type DispatchOptions struct {
	delay int
}

func getDefaultDispatchOptions() *DispatchOptions {
	return &DispatchOptions{
		delay: 0,
	}
}

func DispatchOptionDelay(delayMs int) func(*DispatchOptions) {
	return func(options *DispatchOptions) {
		options.delay = delayMs
	}
}
