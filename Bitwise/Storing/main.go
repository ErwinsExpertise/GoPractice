package main

import (
	"errors"
	"fmt"
	"log"
	"unicode/utf8"
)

const (
	MOV_L   = 1      // A 97
	MOV_R   = 1 << 1 // D 100
	MOV_U   = 1 << 2 // W 119
	MOV_D   = 1 << 3 // S 115
	INVALID = 1 << 5 // Invalid move
)

var CurState uint8

func init() {
	CurState = 0
}
func main() {
	var input string

	for {
		fmt.Scanln(&input)
		key, _ := utf8.DecodeRuneInString(input)
		b, err := Move(key)

		if err != nil {
			log.Println(err)
		}

		fmt.Printf("\nRecieved: %08b", b)
		CurrentState(b)
		fmt.Printf("\nCurrent State: %08b\n", CurState)
		GetDirection()
	}

}

//Move is used to process the key input
func Move(key rune) (uint8, error) {
	if key == 'a' {
		return MOV_L, nil
	}
	if key == 'd' {
		return MOV_R, nil
	}
	if key == 'w' {
		return MOV_U, nil
	}
	if key == 's' {
		return MOV_D, nil
	}
	return INVALID, errors.New("unknown movement")
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
	if (CurState & MOV_L) != 0 {
		fmt.Println("MOV_L")
	}
	if (CurState & MOV_R) != 0 {
		fmt.Println("MOV_R")
	}
	if (CurState & MOV_U) != 0 {
		fmt.Println("MOV_U")
	}
	if (CurState & MOV_D) != 0 {
		fmt.Println("MOV_D")
	}

}
