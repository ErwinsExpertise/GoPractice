package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"unicode/utf8"
)

const (
	MOV_L   = 1      // a 97         00000001
	MOV_R   = 1 << 1 // d 100        00000010
	MOV_U   = 1 << 2 // w 119        00000100
	MOV_D   = 1 << 3 // s 115        00001000
	INVALID = 1 << 5 // Invalid move 00100000
)

var CurState uint8
var pos, size, column int

func init() {
	CurState = 0
	size = int(math.Pow(5, 2))
	pos = (size / 2) + 1
	column = int(math.Sqrt(float64(size)))
}
func main() {
	var input string

	for {
		fmt.Scanln(&input)
		Process(input)
	}
}

func Process(input string) {
	key, _ := utf8.DecodeRuneInString(input)
	b, err := Move(key)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("\nRecieved: %08b", b)
	CurrentState(b)
	fmt.Printf("\nCurrent State: %08b\n", CurState)
	GetDirection()
	CalcDirection()
	ShowDirection()
	fmt.Println()
}

//Move is used to process the key input
func Move(key rune) (uint8, error) {
	switch key {
	case 'a':
		return MOV_L, nil
	case 'd':
		return MOV_R, nil
	case 'w':
		return MOV_U, nil
	case 's':
		return MOV_D, nil
	default:
		return INVALID, errors.New("unknown movement")
	} // end switch(key)

}

//CurrentState returns the current state of
func CurrentState(move uint8) uint8 {
	if move == 1<<5 {
		CurState = 0
		return CurState
	}
	CurState = CurState | move
	return CurState
}

//GetDirection prints the current direction in the bit state
func GetDirection() {
	if hasCurState(MOV_L) {
		fmt.Println("MOV_L")
	}
	if hasCurState(MOV_R) {
		fmt.Println("MOV_R")
	}
	if hasCurState(MOV_U) {
		fmt.Println("MOV_U")
	}
	if hasCurState(MOV_D) {
		fmt.Println("MOV_D")
	}

}

//ShowDirection prints map of where value is
func ShowDirection() {
	for i := 1; i <= size; i++ {
		if i == pos {
			fmt.Print(" + ")
		} else {
			fmt.Print(" * ")
		}
		if i%column == 0 { // checks if the value is multiple
			fmt.Println()
		}
	} //end for

}

//CalcDirection calculates where the value should be on the map
func CalcDirection() {

	if hasCurState(MOV_L) {
		pos--
	}
	if hasCurState(MOV_R) {
		pos++
	}
	if hasCurState(MOV_U) {
		// EXAMPLE (Table size of 25):
		// At 13 need to get to 8
		// 13/5  = 2.6 rounded down to 2 which means theres 2 rows above and it needs to be on row 2
		// Because of POS 13 we know it is as 3
		// 2 * 5 = 10 , 13 - 10 = POS 3
		// [(2-1) * 5] + 3 = 8
		row := int(math.Floor(float64((pos / column))))
		loc := pos - (row * column)

		if row == 1 {
			pos = row * loc
		} else {
			pos = ((row - 1) * column) + loc
		}

	}
	if hasCurState(MOV_D) {
		row := int(math.Ceil(float64((pos / column))))
		loc := pos - (row * column)
		pos = ((row + 1) * column) + loc
	}
}

//hasCurState determines whether or not the bitmask contains a specified state
func hasCurState(mov uint8) bool {
	if (CurState & mov) != 0 { // checks if state exists in the bitmask
		return true
	}
	return false
}
