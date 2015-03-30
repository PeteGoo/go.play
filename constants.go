package main

import "fmt"
import "math"

const s string = "constant" // const declares a constant

func main() {
  fmt.Println(s)

  const n = 500000000

  const d = 3e20 / n    //Constant expressions perform arithmetic with arbitrary precision
  fmt.Println(d)

  fmt.Println(int64(d))   // A numeric constant has no type until it’s given one, such as by an explicit cast.

  /*
  A number can be given a type by using it in a context that requires one,
  such as a variable assignment or function call.
  For example, here math.Sin expects a float64.
  */
  fmt.Println(math.Sin(n))
}
