package main

import (
	"fmt"
	"sync"
	"time"
)

//I realized my "example1" was probably a lot more complex than it needed to be, so this is a
//simpler version that doesn't really show the practical application as much, but is a lot easier
//to digest

//first, initialize the waitgroup
var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	//basically what's going to happen here is we're going to run a goroutine, but because we're not
	//blocking on main, the goroutine won't have a chance to execute all of its code. Running this
	//code in playground, or on your computer, uncomment "wg.Wait()" in order to wait for the
	//goroutine to finish

	wg.Add(1)

	//let's write a silly function that counts to 5. But, slowly
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i + 1)
			time.Sleep(time.Second)
		}
		fmt.Println("Hurrah, I finished counting!")
		wg.Done()
	}()

	//wg.Wait() //UNCOMMENT THIS LINE FOR WAITGROUP MAGIC
	fmt.Println("Program over")
}
