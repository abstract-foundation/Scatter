package main

import (
	"fmt"

	rutils "github.com/Rivative-Foundation/Rutils"
)

func main() {
	fmt.Println("Hello World")
	kp := rutils.NewKeyPair()

	fmt.Println(kp.Address)

	err := rutils.SaveEnvironmentVariable("ADDRESS", "0x3A")

	if err != nil {
		return
	}

	account := rutils.GetEnvironmentVariable("ADDRESS")

	fmt.Println(account)
}
