package main

import (
  "fmt"
  "math"
)

func findArea(number_a, number_b float64) float64 {
  return number_a * number_b * math.Pi
}

func main() {
  number_a := 10.0
  number_b := 20.0

  fmt.Printf("The area is %f\n", findArea(number_a, number_b))
}
