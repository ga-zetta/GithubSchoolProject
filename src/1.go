package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(10)
	b := rand.Intn(10)
	c := a + b
	fmt.Println("The sum of", a, "and", b, "is", c)
}
