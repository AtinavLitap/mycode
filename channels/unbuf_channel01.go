/*
With an unbuffered channel there is no capacity to hold any value before it's received. In this type of channels both a sending and receiving goroutine to be ready at the same instant before any send or receive operation can complete. If the two goroutines aren't ready at the same instant, the channel makes the goroutine that performs its respective send or receive operation first wait.
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)


// this wants a channel to be passed to it
// think of a channel like a 'wormhole'
func CalculateValue(myChan chan int) {
    value := rand.Intn(10)  // create a random value
    fmt.Println("Calculated Random Value: {}", value)
    time.Sleep( time.Second * 10 ) // wait 10 seconds!
    myChan <- value  // send our value back through the wormhole
}

func main() {
    fmt.Println("Go Channel Tutorial")

    // create an empty channel called "values" of type 'int'
    myChan := make(chan int)
    defer close(myChan)

    go CalculateValue(myChan) // pass our channel to our function

    // receiving channels will wait until a value is returned
    value := <-myChan // after 10 seconds a value will return
    fmt.Println(value)
}

