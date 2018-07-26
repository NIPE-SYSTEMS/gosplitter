package gosplitter

import (
	"log"
	"sync"
)

/*
NewSplitter creates a new splitter from a given input channel. It returns an add function that may be used for adding more
outputs to the splitter. The add function returns a new output channel with the given capacity and a remove
function. To prevent memory leaks the remove function must be called when the output channel is no longer needed.

Create splitter:

	input := make(chan interface{})
	add := NewGenericSplitter(input, capacity)

Add output channel:

	output, remove := add()
	defer remove() // ensure that output gets removed
*/
func NewSplitter(input <-chan interface{}, capacity int) func() (<-chan interface{}, func()) {
	add := make(chan chan interface{})
	remove := make(chan chan interface{})
	channels := make(map[chan interface{}]struct{})
	var closed struct {
		sync.Mutex
		b bool
	}
	go func() {
		for {
			select {
			case c := <-add:
				channels[c] = struct{}{}
			case c := <-remove:
				_, ok := channels[c]
				if ok {
					delete(channels, c)
					close(c)
				}
			case v, ok := <-input:
				if !ok {
					closed.Lock()
					closed.b = true
					closed.Unlock()
					for c := range channels {
						delete(channels, c)
						close(c)
					}
					return
				}
				for c := range channels {
					// robust sending to channel (dropping messages for full channels)
					select {
					case c <- v:
					default:
						log.Println("Dropping message, because buffer is full")
					}
				}
			}
		}
	}()
	return func() (<-chan interface{}, func()) {
		c := make(chan interface{}, capacity)
		closed.Lock()
		defer closed.Unlock()
		if closed.b {
			close(c)
		} else {
			add <- c
		}
		return c, func() {
			closed.Lock()
			defer closed.Unlock()
			if !closed.b {
				remove <- c
			}
		}
	}
}
