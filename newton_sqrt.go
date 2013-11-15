package main

import (
    "fmt"
    "math"
)

const delta = 0.0000000000001

func Sqrt(x float64) float64 {
    z := 1.0
    for {
      oldValue := z
      z = z - ( ((z * z) - x) / (2 * z) )
      if math.Abs(oldValue - z) < delta {
        return z
      }
    }
}

func main() {
    fmt.Println(Sqrt(2))
}