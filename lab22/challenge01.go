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

    p1 := Astro{"Andy", 42, "Mars", true}
    p2 := Astro{"Wendy", 62, "Earth", false}
    p3 := Astro{"Aston", 32, "Pluto", true}


    fmt.Println(p1)
    fmt.Println(p2)
    fmt.Println(p3)


}
