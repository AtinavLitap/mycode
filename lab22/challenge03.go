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


    type nasaMission struct {
        people []Astro
        number int
        message string
    }

    p := []Astro {
        {"Andy", 42, "Mars", true},
        {"Wendy", 62, "Earth", false},
        {"Aston", 32, "Pluto", true},
    }
    
    s := nasaMission { p, 3, "success"}



    fmt.Println(s)


}
