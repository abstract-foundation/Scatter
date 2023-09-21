package main

import (
	"fmt"

	rutils "github.com/Awesome-Sauces/Rutils"
)

func main() {
	fmt.Println("Hello World")
	kp := rutils.NewKeyPair()

	fmt.Println(kp.Address)
}
