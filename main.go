package main 

import (
    "fmt"
    "os"
    "io/ioutil"
    u "github.com/ardeshir/version"
)

type point struct{ x, y int64 }

func (p point) String() string {
    return fmt.Sprintf("Point(%d, %d)", p.x, p.y)
}

var ( 
 debug bool = true
 version string = "0.0.1"
)

var (
    
    i = 100
    f = 3.14
    b = true
    s = "Clear is better than clever."
    p = point{256, 101}
    
    )
    
var bi int64  = -922337203685
var ui uint64 =  1844679551625

func outputStdout() {
    fmt.Println(p)
    fmt.Printf("%v\n", p)
    /* fmt.Println(i, f, b, bi, ui, p)
    fmt.Printf(" i = %v\n f = %v\n b = %v\n s = %v\n bi = %v\n ui = %v\n p = %v\n", i,f,b,s,bi,ui,p)
    fmt.Println("\n")
    fmt.Printf(" i = %#v\n f = %#v\n b = %#v\n s = %#v\n bi = %#v\n ui = %#v\n p = %#v\n", i,f,b,s,bi,ui,p)
    fmt.Println("\n")
    fmt.Printf(" i = %T\n f = %T\n b = %T\n s = %T\n bi = %T\n ui = %T\n p = %T\n", i,f,b,s,bi,ui,p) */
    //fmt.Printf("\n", i,f,b,s,bi,ui,p)
    
}

func outputWriter() (string, error) {
    file , err := ioutil.TempFile("", "tmp.")
    if err != nil {
        return "", err
    }
    
    defer file.Close()
    
    fmt.Fprintf(file, " i = %#v\n f = %#v\n b = %#v\n s = %#v\n bi = %#v\n ui = %#v\n p = %#v\n", i,f,b,s,bi,ui,p)
    return file.Name(), nil
}

func main() {

    outputStdout();
    
    n, err := outputWriter();
    if err != nil {
        fmt.Printf("Error occured: %s\n", err)
        os.Exit(-1)
    }
    
    fmt.Printf("Filename: %s\n", n)
    //  end of main -- helper functions
      if debugTrue() {
        u.V(version)
      }
}

// if debug true func

// Function to check env variable DEFAULT_DEBUG bool
func debugTrue() bool {
    
     if os.Getenv("DEFAULT_DEBUG") != "" {
        return true
     }  
     return false 
}
