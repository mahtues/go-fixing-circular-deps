package concrete

import (
	"github.com/mahtues/go-fixing-circular-deps/internal/a"
)

type BStruct struct {
	a a.AInterface
}

func New() *BStruct {
	return &BStruct{}
}

func (b *BStruct) Inject(a a.AInterface) {
	b.a = a
}

func (b *BStruct) Bar() string {
	return "bar from b"
}

func (b *BStruct) Bnm() string {
	return "bnm from b: " + b.a.Foo()
}
