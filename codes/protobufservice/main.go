package main

import (
	"time"
)

func main() {
	go doServerAction()
	time.Sleep(time.Second)
	doClientAction()
}
