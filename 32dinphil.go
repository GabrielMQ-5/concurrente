package main

import (
	"fmt"
	"sync"
	"time"
)

var fork []sync.Mutex
var fin chan bool

const n = 5

func philosopher(id int, nombre string) {
	izq := id
	der := (id + 1) % n
	for i := 0; i < 10; i++ {
		fmt.Printf("%s está pensando\n", nombre)
		fork[izq].Lock()
		time.Sleep(time.Millisecond)
		fork[der].Lock()
		fmt.Printf("%s está comiendo\n", nombre)
		fork[izq].Unlock()
		fork[der].Unlock()
	}
	fin <- true
}

func main() {
	fork = make([]sync.Mutex, n)
	philosophers := []string{"Platón", "Sócrates", "Descartes", "Nietzsche", "Schopenhauer"}
	for i := 0; i < n; i++ {
		go philosopher(i, philosophers[i])
	}
	for i := 0; i < n; i++ {
		<-fin
	}
}
