package main

import (
	"fmt"
	"time"
)

func newChan(deadline time.Duration) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		time.Sleep(deadline)
	}()
	return ch
}

func solver(channels ...<-chan interface{}) <-chan interface{} {

	out := make(chan interface{})
	done := make(chan interface{})

	for i := range channels {
		go func(ch <-chan interface{}) {
			select {
			case tmp := <-ch:
				close(done)
				out <- tmp
			case <-done:
				return
			}
		}(channels[i])
	}
	<-done

	return out
}

func main() {
	start := time.Now()

	<-solver(
		newChan(2*time.Hour),
		newChan(5*time.Minute),
		newChan(36000*time.Millisecond),
		newChan(3*time.Second),
		newChan(5*time.Second),
		newChan(1*time.Hour),
		newChan(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
