package main

import (
  "fmt"
  "math"
)

func findArea(number_a, number_b float64) float64 {
  return number_a * number_b * math.Pi
}

func main() {
  var number_a float64 = 10
  var number_b float64 = 20

  fmt.Printf("The area is %f\n", findArea(number_a, number_b))
}
