package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func receive(chRecv <-chan bool) {
	fmt.Println("Try to read from Channel.")
	rslt := <-chRecv
	fmt.Printf("Read from Channel - %v, %p: %v\n", chRecv, chRecv, rslt)
}

func send(chSend chan<- bool) {
	fmt.Println("Try to send to Channel.")
	chSend <- true
	fmt.Println("Sent to Channel.")
}

func TestChanelExceptions(t *testing.T) {
	ch := make(chan bool)
	go func(ch <-chan bool) {
		receive(ch)
	}(ch)
	send(ch)

	fmt.Println("Test basic done.")

	time.Sleep(time.Second)

	fmt.Println("******************\nTest close chanel while recving from it.")
	go func(ch chan bool) {
		runtime.Gosched()
		time.Sleep(time.Second)
		close(ch)
		fmt.Println("Channel Closed.")
	}(ch)
	receive(ch)

	fmt.Println("******************\nTest send to a closed chanel")
	ch = make(chan bool)
	go func(ch <-chan bool) {
		receive(ch)
	}(ch)
	close(ch)
	defer recoverable()
	send(ch)
}

func recoverable() {

	r := recover()

	if r != nil {
		fmt.Println("recovered from ", r)
		time.Sleep(time.Second)
		fmt.Println("All test done.")
	} else {
		panic("Test Error, failed to catch the expected panic.")
	}
}
