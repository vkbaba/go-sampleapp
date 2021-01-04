package main

import (
	"fmt"
	"github.com/vkbaba/go-sampleapp/math"
)

func main() {
    sum := math.Add(3, 4)
	fmt.Printf("Hello,\"%d\"!", sum)
}
