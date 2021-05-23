package main

import (
	"fmt"
	"sync"
)

/*In this code, we'll explore the basic useage of wait groups in go. Wait groups are just what they
sound like - they tell a goroutine (often your main() function), to wait and not return until other
goroutines are done*/

//first, we need to create a wait group. This will keep track of all the functions that are using
//a given channel. To keep this example basic, we'll only use one channel and one wait group
var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	//let's create a channel to use with our waitgroup
	c := make(chan int)

	fmt.Println("Listening now")

	//Let's start a goroutine that will listen on our channel. The function is defined below main
	go listen(c)

	//since we'll be running our function twice, we'll add 2 to the wait group (if we were running
	//3 times, we would add 3, etc)
	wg.Add(2)

	fmt.Println("Now we're starting our functions")

	//now, lets send c to a few goroutines using a function we'll define below
	go myFunc(c, wg)
	go myFunc(c, wg)

	fmt.Println("And now, we wait")

	//Now that our goroutines are running, we want to block on main, so that it doesn't return before
	//our goroutines are finished doing their work. We'll do that with wg.Wait(), which blocks until
	//it gets all its done() signals from our goroutines
	wg.Wait()

	//waitgroup is done waiting! Let's close our channel
	close(c)

	fmt.Println("Channel is closed! Good work everyone, we've avoided deadlock.")
}

//Let's go ahead and define our function. It will take 1 channel and 1 wait group as arguments
func myFunc(c chan int, wg *sync.WaitGroup) {
	//we'll just execute some silly code so that this function has something to send to channel c
	for i := 1; i < 10; i++ {
		c <- i * i
	}
	//our code has been executed, but we can't just close the channel, because another goroutine
	//might still be using it, so that would give us an error. Instead, we'll just tell the wait
	//group that *this* goroutine is done using our channel
	fmt.Println("All done! Sending done to my wg")
	wg.Done()
	return //I often explicitly return when I don't technically need to. It's just for clarity, really
}

//This just listens to our channel and prints out data from it until the channel is closed
func listen(c chan int) {
	//We'll just have a for loop continually listening on our channel
	for {
		val, ok := <-c  //checks to see if we're receing a signal from channel c
		if ok != true { //if we aren't
			return //end the goroutine
		} else {
			fmt.Println(val) //but if we are, we'll just print out what we're receiving
		}
	}
}
