package main

import (
	"flag"
	"fmt"
	"log"
)

var op string
var num1, num2 int

func init() {
	flag.StringVar(&op, "op", "", "Specify which bitwise operation to perform( xor , and , or , sl , sr")
	flag.IntVar(&num1, "num1", 1, "Bit 1")
	flag.IntVar(&num2, "num2", 2, "Bit 2")

	flag.Parse()
}

func main() {
	n1, n2 := uint8(num1), uint8(num2)

	switch op {
	case "xor":
		XORFunc(n1, n2, op)
		break
	case "and":
		ANDFunc(n1, n2, op)
		break
	case "or":
		ORFunc(n1, n2, op)
		break
	case "sl":
		ShiftLFunc(n1, n2, op)
		break
	case "sr":
		ShiftRFunc(n1, n2, op)
		break
	default:
		log.Panicln("Invalid operation or none specified")
		break

	}
	fmt.Println()
}

//XORFunc performs ^ bitwise operations
func XORFunc(n1, n2 uint8, fname string) {
	fmt.Printf("\n%s returned: %08b", fname, n1^n2)
}

//ANDFunc performs & bitwise operation
func ANDFunc(n1, n2 uint8, fname string) {
	fmt.Printf("\n%s returned: %08b", fname, n1&n2)
}

//ORFunc performs | bitwise operation
func ORFunc(n1, n2 uint8, fname string) {
	fmt.Printf("\n%s returned: %08b", fname, n1|n2)
}

//ShiftLFunc performs << bitwise operation
func ShiftLFunc(n1, n2 uint8, fname string) {
	fmt.Printf("\n%s returned: %08b", fname, n1<<n2)
}

//ShiftRFunc performs >> bitwise operation
func ShiftRFunc(n1, n2 uint8, fname string) {
	fmt.Printf("\n%s returned: %08b", fname, n1>>n2)
}
