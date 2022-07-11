package hermessenger

import (
	"testing"
)

func TestDispatchOptionsDelay(t *testing.T) {
	dispatchOptions := getDefaultDispatchOptions()

	delayMs := 1234
	optionFn := DispatchOptionDelay(delayMs)
	optionFn(dispatchOptions)
	if dispatchOptions.delay != delayMs {
		t.Fatalf("got %d, wanted %d", dispatchOptions.delay, delayMs)
	}
}
