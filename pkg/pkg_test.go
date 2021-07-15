package pkg_test

import (
	"fmt"
	"reflect"
	"runtime"
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
	fmt.Printf("%p %p\n", a, b)
	sum := 0
	for i := 0; i < 10000; i++ {
		c := make([]byte, 1024*1024)
		sum += len(c)
	}
	fmt.Println(sum)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		runtime.GC()
		wg.Done()
	}()
	go func() {
		b = b[1:2]
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("%p %p\n", a, b)
	return b
}

func TestSomething(t *testing.T) {
	str := "123"
	a := Convert([]byte(str))
	b := Copy(a)
	Mid(a, b.S)
}
