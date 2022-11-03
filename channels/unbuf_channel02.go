/* Alta3 Research | RZFeeser
   Channels - A simple example.
   Think of a channel like a wormhole. Anything placed into it,
   is instantly moved between goroutines.*/


/* This example creates another unbuffered channel. This example attempts to demonstrate some of the limitations of this channel type.  */

/* Running this example will result in a single instance of, "I run after the data is added to the channel". That's because our unbuffered channel is full. Until it is read out, all other functions must wait.

The solution is a buffered channel. Buffered channels allow queue in which to drop information, and keep running.*/


package main

import (
    "fmt"
    "math/rand"
    "time"
)


// this wants a channel to be passed to it
// think of a channel like a 'wormhole'
func CalculateValue(values chan int) {
    value := rand.Intn(10)  // create a random value
    fmt.Println("Calculated Random Value: {}", value)
    time.Sleep( time.Second * 10 ) // wait 10 seconds!
    values <- value  // send our value back through the wormhole
    // after the data is through the wormhole
    fmt.Println("I run after the data is added to the channel")
}

func main() {
    fmt.Println("Go Channel Tutorial")

    // create an empty channel called "values" (UNBUFFERED)
    values := make(chan int)
    defer close(values)

    go CalculateValue(values) // pass our channel to our function
    go CalculateValue(values)
    go CalculateValue(values)
    go CalculateValue(values)

    // wait until a value is returned
    value := <-values // after 10 seconds a value will return

    time.Sleep( time.Second ) // wait a single second (after data is read from the channel)
    fmt.Println(value) // I run after the data is read from the channel
}