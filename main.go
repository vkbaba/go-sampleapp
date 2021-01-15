package main

import (
	"fmt"
	"github.com/vkbaba/go-sampleapp/math"
)

func main() {
    a, b := 3, 4
	sum := math.Add(a, b)
	fmt.Printf("Hello, %d + %d = %d ???\n", a, b, sum)
}
