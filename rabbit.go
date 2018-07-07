package main

import (
	"fmt"
	"math/rand"
	"time"
)

var stop chan bool
var run [2]chan int
var winner chan string
var running int

const length = 100

func rabbit() {
	distance := 0
	distance_run := 0
	for {
		if distance_run >= length {
			break
		}
		distance_run = <-run[0]
		distance = getRandomDistance()
		time.Sleep(50 * time.Millisecond)
		distance_run += distance
		run[0] <- distance_run
	}
	running--
}

func turtle() {
	distance := 1
	distance_run := 0
	for distance_run < length {
		distance_run = <-run[1]
		time.Sleep(50 * time.Millisecond)
		distance_run += distance
		run[1] <- distance_run
	}
	running--
}

func watcher() {
	d_rabbit, d_turtle := 0, 0
	for running > 0 {
		d_rabbit = <-run[0]
		d_turtle = <-run[1]
		d_rabbit = bite(d_rabbit, d_turtle)
		if d_rabbit == 100 {
			select {
			case winner <- "\nHare won\n":
			default:
			}
		} else if d_turtle == 100 {
			select {
			case winner <- "\nTortoise won\n":
			default:
			}
		}
		fmt.Printf("Tortoise at %3d, Hare at %3d\n", d_turtle, d_rabbit)
		run[0] <- d_rabbit
		run[1] <- d_turtle
	}
	fmt.Printf("%s\n", <-winner)
	stop <- true
}

func bite(dr, dt int) int {
	if dr > 100 {
		dr = 100
	}
	if dr == dt && dr != 100 {
		dr -= 4
		if dr < 0 {
			dr = 0
		}
		fmt.Printf("\nTortoise has bit Hare at %3d CRUNCH\n\n", dt)
	}
	return dr
}

func getRandomDistance() int {
	random := rand.Intn(10)
	random /= rand.Intn(3) + 2
	return random
}

func main() {
	running = 2
	stop = make(chan bool, 1)
	run[0] = make(chan int, 1)
	run[1] = make(chan int, 1)
	winner = make(chan string, 1)
	run[0] <- 0
	run[1] <- 0
	go watcher()
	go rabbit()
	go turtle()
	<-stop

}
