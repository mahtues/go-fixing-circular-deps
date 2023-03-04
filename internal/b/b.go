package b

import (
	"github.com/mahtues/go-fixing-circular-deps/internal/a/xa"
)

type B struct {
	a xa.Xa
}

func New() *B {
	return &B{}
}

func (b *B) Inject(a xa.Xa) {
	b.a = a
}

func (b *B) Bar() string {
	return "bar from b"
}

func (b *B) Bnm() string {
	return "bnm from b: " + b.a.Foo()
}
