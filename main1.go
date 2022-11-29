package main

import (
	"fmt"
	"jonathangersam.com/embed_tests/internal/animportedpkg"
)

func main() {
	animportedpkg.SayHi()
	fmt.Println("main ends")
}
