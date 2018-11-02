package main 

import "fmt"

var (
    
    i = 100
    f = 3.14
    b = true
    s = "Clear is better than clever."
    p = struct{ x, y int64 }{256, 101}
    
    )
    
var bi int64  = -922337203685
var ui uint64 =  1844679551625

func main() {
    // fmt.Println(i, f, b, bi, ui, p)
    fmt.Printf(" i = %v\n f = %v\n b = %v\n s = %v\n bi = %v\n ui = %v\n p = %v\n", i,f,b,s,bi,ui,p)
    fmt.Println("\n")
    fmt.Printf(" i = %#v\n f = %#v\n b = %#v\n s = %#v\n bi = %#v\n ui = %#v\n p = %#v\n", i,f,b,s,bi,ui,p)
    fmt.Println("\n")
    fmt.Printf(" i = %T\n f = %T\n b = %T\n s = %T\n bi = %T\n ui = %T\n p = %T\n", i,f,b,s,bi,ui,p)
    //fmt.Printf("\n", i,f,b,s,bi,ui,p)
}