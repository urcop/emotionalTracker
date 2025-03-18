package main

import (
	"github.com/urcop/emotionalTracker/app"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		app.NewHttpServer().Start()
	}()

	go func() {
		defer wg.Done()
		app.NewGRPCServer().Start()
	}()
	wg.Wait()
}
