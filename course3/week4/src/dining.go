package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
}

type Philosopher struct {
	num int
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(Chopstick)
	},
}

var w sync.WaitGroup
// var wg sync.
func (p Philosopher) eat(ch chan int) {
	
	<-ch

	fmt.Printf("starting to eat %d\n", p.num)
	// pick up chopsticks in any order
	left := pool.Get()
	right := pool.Get()
	// then return the chopsticks
	pool.Put(left)
	pool.Put(right)
	fmt.Printf("finishing eating %d\n", p.num)

	ch <- 1
	w.Done()
}

func main() {
	// ch allows no more than 2 philosophers to eat concurrently
	ch := make(chan int, 2)

	// initialize chopsticks pool
	for i := 0; i < 5; i++ {
		pool.Put(new(Chopstick))
	}

	// initialize philosophers
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		// each philosopher is numbered, 1 through 5
		philosophers[i] = &Philosopher{i + 1}
	}

	// let them start
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ { // only eat 3 times
			w.Add(1)
			go philosophers[i].eat(ch)
		}
	}

	ch <- 1
	ch <- 1

	w.Wait()
}