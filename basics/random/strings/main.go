package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	pointer uintptr
	lenght  int
}

func main() {
	empty := ""
	dump(empty) // "": {pointer:0 lenght:0}

	hello := "hello"
	dump(hello)     // "hello": {pointer:33405286 lenght:5}
	dump("hello")   // "hello": {pointer:33405286 lenght:5}
	dump("hello!")  // "hello!": {pointer:33405556 lenght:6}
	dump(hello[:1]) // "h": {pointer:61737318 lenght:1}
	dump(hello[1:]) // "ello": {pointer:61737319 lenght:4}
}

func dump(s string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}
