package main

import (
	"sip/app"
	"sync"
)

var wg sync.WaitGroup

func main() {
	app.Run()
}
