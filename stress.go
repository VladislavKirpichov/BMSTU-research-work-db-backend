package main

import (
	"net/http"
	"sync"
)

func helloWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	http.Get("http://localhost:8080/api/users")
}

func callServer(wg *sync.WaitGroup) {
	defer wg.Done()
	http.Get("http://localhost:8080/api/users")
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go callServer(wg)
	}

	wg.Wait()
}
