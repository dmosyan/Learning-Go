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

// 👉 %s prints []byte as string
// 👉 %c prints a rune as a character
// 👉 %x prints the bytes of a string in hexadecimal notation
// 👉 % x like %x but prints the bytes with spaces
// 👉 for range jumps over the runes in a string
// 👉 aString[index] returns a byte
// 👉 aString[start:stop] returns a string
// 👉 rune(aString) returns a rune slice by putting the runes inside the string
// 👉 string(anInteger) returns a string by utf-8 encoding the integer
// 👉 byteSlice = append(byteSlice, aString...)
// You can append a string to a byte slice
// 👉 utf8 and unicode packages contain helper functions for working with runes
