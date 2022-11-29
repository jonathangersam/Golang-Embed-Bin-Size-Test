package animportedpkg

import (
	_ "embed"
	"testing"
)

////go:embed "fixtures/some_huge_embed.bin"
//var someHugeEmbed []byte

func TestSomething(t *testing.T) {
	t.Log("test ran")
}
