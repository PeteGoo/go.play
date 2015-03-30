package main

import "fmt"
import "time"

func main() {
  i := 2
  fmt.Println("write ", i, " as ")
  switch i {
    case 1:
      fmt.Println("one")
    case 2:
      fmt.Println("two")
    case 3:
      fmt.Println("three")
  }

  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
      fmt.Println("it's the weekend")
    default:
      fmt.Println("it's a weekday")
  }

  // You can use commas to separate multiple expressions in the same case statement.
  // We use the optional default case in this example as well.


  t := time.Now()
  switch {
    case t.Hour() < 12:
      fmt.Println("it's before noon")
    default:
      fmt.Println("it's after noon")
  }

  // switch without an expression is an alternate way to express if/else logic.
  // Here we also show how the case expressions can be non-constants.
}
