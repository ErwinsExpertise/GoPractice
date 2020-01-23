package main

import "fmt"

const (
	//kb = 1024
	//mb = 1024 * kb
	//gb = 1024 * mb
	_  = iota
	kb = 1 << (iota * 10) // 1 * 10
	mb = 1 << (iota * 10) // 2 * 10
	gb = 1 << (iota * 10) // 3 * 10
)

func main() {
	fmt.Printf("&d\t\t%b\n", kb, kb)
	fmt.Printf("&d\t\t%b\n", mb, mb)
	fmt.Printf("&d\t%b\n", gb, gb)
}
