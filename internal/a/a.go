package a

import (
	"github.com/mahtues/go-fixing-circular-deps/internal/b/xb"
)

type A struct {
	b xb.Xb
}

func New() *A {
	return &A{}
}

func (a *A) Inject(b xb.Xb) {
	a.b = b
}

func (a *A) Foo() string {
	return "foo from a"
}

func (a *A) Asd() string {
	return "asd from a: " + a.b.Bar()
}
