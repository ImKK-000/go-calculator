package main

import (
  "fmt"
  "math"
)

func findArea(number_a, number_b float64) float64 {
  return number_a * number_b * math.Pi
}

func main() {
  var number_a float64
  var number_b float64

  fmt.Scanf("%d %d", &number_a, &number_b)
  fmt.Println(findArea(number_a, number_b))
}
