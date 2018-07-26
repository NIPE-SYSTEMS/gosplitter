package main

import (
	"sync"
	"testing"
)

func TestSplitter(t *testing.T) {
	c := make(chan interface{})
	add := NewSplitter(c, 3)
	var wgAdded sync.WaitGroup
	var wgGoroutines sync.WaitGroup
	for i := 0; i < 10; i++ {
		wgAdded.Add(1)
		wgGoroutines.Add(1)
		go func(i int) {
			defer wgGoroutines.Done()
			c, remove := add()
			defer remove()
			wgAdded.Done()
			for v := range c {
				t.Log(v)
			}
		}(i)
	}
	wgAdded.Wait()
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
	wgGoroutines.Wait()
}
