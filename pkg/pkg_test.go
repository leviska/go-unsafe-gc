package pkg_test

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"unsafe"
)

func StringToByteUnsafe(s string) []byte {
	strh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	var sh reflect.SliceHeader
	sh.Data = strh.Data
	sh.Len = strh.Len
	sh.Cap = strh.Len
	return *(*[]byte)(unsafe.Pointer(&sh))
}

func Convert(s []byte) []byte {
	return StringToByteUnsafe(string(s))
}

type T struct {
	S []byte
}

func Copy(s []byte) T {
	return T{S: s}
}

func Mid(a []byte, b []byte) []byte {
	fmt.Printf("%p %s %p %s\n", a, a, b, b)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		b = b[1:2]
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("%p %s %p %s\n", a, a, b, b)
	return b
}

func TestSomething(t *testing.T) {
	str := "123"
	a := Convert([]byte(str))
	b := Copy(a)
	Mid(a, b.S)
}
