package main

import (
    "fmt"
    "os"
    "io/ioutil"
  //  "strings"
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


func main() {
    fmt.Println("\n\n")

    scanWithSscan();
    scanWithSscanf()
    scanWithSscanln()

    outputStdout();

    n, err := outputWriter();
    if err != nil {
        fmt.Printf("Error occured: %s\n", err)
        os.Exit(-1)
    }
    fmt.Println("\n")
    fmt.Printf("Filename: %s\n", n)
    //  end of main -- helper functions
      if debugTrue() {
        u.V(version)
      }
}

func scanWithSscanln() {
  fmt.Println("## scanWithSscan ##")
  var d1, d2 int
  var s1 string
  fmt.Printf("Before scanning with Sscan: %d, %d, %s\n", d1, d2, s1)
  if _, err := fmt.Sscanln("5 7 9\n", &d1, &d2, &s1); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("After scanning with Sscanln: %d, %d, %s\n", d1, d2, s1)
  fmt.Println("\n")
}

func scanWithSscanf() {
  fmt.Println("## scanWithSscanf ##")
  var d1, d2 int
  var s1 string
  fmt.Printf("Before scanning with Sscan: %d, %d, %s\n", d1, d2, s1)
  if _, err := fmt.Sscanf("5, 7 and 9 are my values", "%d, %d and %s", &d1, &d2, &s1); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("After scanning with Sscanf: %d, %d, %s\n", d1, d2, s1)
  fmt.Println("\n")
}


func scanWithSscan() {
  fmt.Println("## scanWithSscan ##")
  var d1, d2 int
  var s1 string
  fmt.Printf("Before scanning with Sscan: %d, %d, %s\n", d1, d2, s1)
  if _, err := fmt.Sscan("5 7\n9", &d1, &d2, &s1); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("After scanning with Sscan: %d, %d, %s\n", d1, d2, s1)
  fmt.Println("\n")
}

// fmt.Printf stuff
func outputStdout() {
    fmt.Println("\n\n")
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

// if debug true func
// Function to check env variable DEFAULT_DEBUG bool
func debugTrue() bool {

     if os.Getenv("DEFAULT_DEBUG") != "" {
        return true
     }
     return false
}
