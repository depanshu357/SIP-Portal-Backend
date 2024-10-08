package main

import (
	"sip/app"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go app.Run(&wg)
	wg.Wait()
}
