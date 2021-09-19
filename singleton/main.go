package main

import (
	"fmt"
	"sync"
	"time"
)

type lolInKek int64

type someKek struct {
	mu    sync.Mutex
	value *lolInKek
}

func (s *someKek) getSomeLol() *lolInKek {
	if s.value == nil {
		s.mu.Lock()
		defer s.mu.Unlock()

		if s.value == nil {
			value := lolInKek(10)
			s.value = &value
			fmt.Println("some lol created")
			return s.value
		}
		fmt.Println("some lol is already created")
		return s.value
	}
	fmt.Println("some lol is already created")

	return s.value
}

func main() {

	kek := &someKek{}

	for i := 0; i < 10; i++ {
		go kek.getSomeLol()
	}

	time.Sleep(time.Second)
}
