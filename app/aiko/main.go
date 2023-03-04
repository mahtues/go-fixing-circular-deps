package main

import (
	"fmt"

	"github.com/mahtues/go-fixing-circular-deps/internal/a"
	"github.com/mahtues/go-fixing-circular-deps/internal/b"
)

func main() {
	asd := a.New()
	fmt.Println(asd.Foo())

	bnm := b.New()
	fmt.Println(bnm.Bar())

	asd.Inject(bnm)
	fmt.Println(asd.Asd())

	bnm.Inject(asd)
	fmt.Println(bnm.Bnm())
}
