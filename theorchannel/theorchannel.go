package theorchannel

import (
	"fmt"
	"time"
)

/*
The OR channel:
1 channel completed, trigger closing all channels
*/
func TheORChannelExperiment() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(sig(2*time.Hour), sig(5*time.Minute), sig(2*time.Second), sig(1*time.Hour), sig(1*time.Minute))
	fmt.Printf("Done after: %v\n", time.Since(start))

}

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	done := make(chan interface{})
	go func() {
		defer close(done)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], done)...):
			}
		}
	}()
	return done
}
