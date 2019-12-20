package main

import (
	"errors"
	"fmt"
	"log"
	"math"
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
var pos, size int

func init() {
	CurState = 0
	size = int(math.Pow(5, 2))
	pos = (size / 2) + 1
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
		CalcDirection()
		ShowDirection()
		fmt.Println()
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

//ShowDirection prints map of where value is
func ShowDirection() {
	columns := math.Sqrt(float64(size))
	for i := 1; i <= size; i++ {
		if i == pos {
			fmt.Print(" A ")
		} else {
			fmt.Print(" * ")
		}
		if i%int(columns) == 0 { // checks if the value is multiple
			fmt.Println()
		}
	} //end for

}

//CalcDirection calculates where the value should be on the map
func CalcDirection() {
	column := int(math.Sqrt(float64(size)))

	if (CurState & MOV_L) != 0 {
		pos--
	}
	if (CurState & MOV_R) != 0 {
		pos++
	}
	if (CurState & MOV_U) != 0 {
		// EXAMPLE (Table size of 25):
		// At 13 need to get to 8
		// 13/5  = 2.6 rounded down to 2 which means theres 2 rows left and it needs to be on row 2
		// Because of POS 13 we know it is as 3
		// 2 * 5 = 10 , 13 - 10 = POS 3
		// [(2-1) * 5] = 5 , 5 + 3 = 8
		row := int(math.Floor(float64((pos / column))))
		loc := pos - (row * column)

		if row == 1 {
			pos = row * loc
		} else {
			pos = ((row - 1) * column) + loc
		}

	}
	if (CurState & MOV_D) != 0 {
		row := int(math.Ceil(float64((pos / column))))
		loc := pos - (row * column)
		pos = ((row + 1) * column) + loc
	}
}
