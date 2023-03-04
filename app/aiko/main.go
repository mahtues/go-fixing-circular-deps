package main

import (
	"fmt"

	a "github.com/mahtues/go-fixing-circular-deps/internal/a/concrete"
	b "github.com/mahtues/go-fixing-circular-deps/internal/b/concrete"
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
