package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Use all of the cores
// We dont need this anymore, as default after g01.5 we use all cores
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 100; i < 106; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for b := 0; b < 45; b++ {
		fmt.Println("Bar:", b)
		time.Sleep(time.Duration(2 * time.Millisecond))
	}
	wg.Done()
}
