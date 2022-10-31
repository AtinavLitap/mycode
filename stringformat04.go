package main

import (
    "fmt"
)


func main() {

    uri := "https://example.org:6001/v2/snacks?"
    r := "req=snickers"
    q := "quantity=1"
    s := "size=king"

    result := uri+r+"&"+q+"&"+s

    fmt.Println(result)
    //fmt.Println(uri,r,q,s)          # adds spaces between variables

    // Another way to do this is:
    const uri1 = "https://example.org:6001/v2/snacks?"
    var r1, q1, s1 string = "req=snickers", "quantity=1", "size=king"
    // create the string, but don't render it yet
    res := fmt.Sprintf("%s%s&%s&%s", uri1, r1, q1, s1)
    fmt.Println(res) // finished URI

}
