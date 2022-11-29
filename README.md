# Golang Embed Bin Size Test

WHY
* Demonstration purposes.

WHAT
* Tests whether an embed not used in main will increase final bin size.

USAGE
1. run `./build.sh`
2. run `ls -l` to see the output file sizes. 
* `main1` imports a package that references a large embed, but never executes it
* `main2` imports a package that references a large embed, and executes it.
3. notice `main1` is much smaller than `main2`

CONCLUSION: 
* embeds in tests are **NOT INCLUDED** the binary
* embeds in an imported pkg but referenced by the importer (e.g. main) are **NOT INCLUDED**
* embeds imported but could never execute are **NOT INCLUDED**

```
//go:embed "fixtures/some_huge_embed.bin"
var someHugeEmbed []byte

func SayHi() {
	fmt.Println("hi")

	aReference := someHugeEmbed // reference made, but never used
	if false {
		fmt.Println(aReference) // can never run
	}
}
```

* embeds imported and actually used are **INCLUDED**

```
//go:embed "fixtures/some_huge_embed.bin"
var someHugeEmbed []byte

func SayHello() {
	fmt.Println("hello")

	aReference := someHugeEmbed
	if true {
		fmt.Println(aReference) // can run
	}
}
```
