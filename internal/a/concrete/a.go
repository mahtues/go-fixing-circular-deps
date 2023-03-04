package concrete

import (
	"github.com/mahtues/go-fixing-circular-deps/internal/b"
)

type AStruct struct {
	b b.BInterface
}

func New() *AStruct {
	return &AStruct{}
}

func (a *AStruct) Inject(b b.BInterface) {
	a.b = b
}

func (a *AStruct) Foo() string {
	return "foo from a"
}

func (a *AStruct) Asd() string {
	return "asd from a: " + a.b.Bar()
}
