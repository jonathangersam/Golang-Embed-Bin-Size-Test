package animportedpkg

import (
	_ "embed"
	"fmt"
)

//go:embed "fixtures/some_huge_embed.bin"
var someHugeEmbed []byte

func SayHi() {
	fmt.Println("hi")

	aReference := someHugeEmbed
	if false {
		fmt.Println(aReference)
	}
}

func SayHello() {
	fmt.Println("hello")

	aReference := someHugeEmbed
	if true {
		fmt.Println(aReference)
	}
}
