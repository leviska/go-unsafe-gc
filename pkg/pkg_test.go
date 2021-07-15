package pkg_test

import (
	"testing"

	"github.com/leviska/go-unsafe-gc/other"
	"github.com/leviska/go-unsafe-gc/pkg"
)

func TestSomething(t *testing.T) {
	str := "123"
	a := pkg.Convert([]byte(str))
	b := other.Copy(a)
	pkg.Mid(a, b.S)
}
