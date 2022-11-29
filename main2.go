package main

import (
	"fmt"
	"jonathangersam.com/embed_tests/internal/animportedpkg"
)

func main() {
	animportedpkg.SayHello()
	fmt.Println("main 2 ends")
}
