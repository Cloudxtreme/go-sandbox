package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "hello, world!"
	vec := *(*[]byte)(unsafe.Pointer(&str))
	fmt.Println(vec)
}
