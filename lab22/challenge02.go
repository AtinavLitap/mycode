package main

import (
    "fmt"
)

func main() {

    type Astro struct {
        name string
        age int
        mission string
        isneeded bool
    }

    p := []Astro {

        {"Andy", 42, "Mars", true},
        {"Wendy", 62, "Earth", false},
        {"Aston", 32, "Pluto", true},

    }

    fmt.Println(p)


}
