package main

import "fmt"

func main(){
  var a string = "The string" // var defines a variable
  fmt.Println(a)

  var b, c int = 1,2 // You can define more than one at a time
  fmt.Println(b, c)

  var d = true // Go wil infer types
  fmt.Println(d)

  var e int // Go associates a zero-value with an unassigned variable
  fmt.Println(e)

  f:= "short" // The := syntax is shorthand for declare and initialize
  fmt.Println(f)

}
