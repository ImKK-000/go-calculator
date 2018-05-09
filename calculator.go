package main

import (
  "fmt"
  "math"
)

func main() {
  var number_a float64
  var number_b float64

  fmt.Scanf("%d %d", &number_a, &number_b)
  fmt.Println(number_a * number_b * math.Pi)
}
