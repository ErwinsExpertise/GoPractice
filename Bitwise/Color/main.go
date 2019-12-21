package main

import "fmt"

//RGB
//Dodger Blue (30,144,255)
var (
	R = 30
	G = 144
	B = 255
)

func main() {
	// Color has an 8-bit stream

	color := R << 16
	fmt.Printf("\nColor: %b", color)
	color += G << 8
	fmt.Printf("\nColor: %b", color)
	color += B
	fmt.Printf("\nColor: %b", color)

	mask := 255 // 0xFF

	r := (color >> 16) & mask
	g := (color >> 8) & mask
	b := color & mask

	fmt.Printf("\n\nr: %b\ng: %b\nb: %b\n", r, g, b)
	fmt.Printf("\n\nDecimals\nr: %d\ng: %d\nb: %d\n", r, g, b)

	// Darken color by 50% easy way
	// To lighten by 50% just use 0x7f7f7f

	color = (color & 0xfefefe) >> 1

	r = (color >> 16) & mask
	g = (color >> 8) & mask
	b = color & mask

	fmt.Println("\nDarkened by 50%")
	fmt.Printf("\nColor: %b", color)
	fmt.Printf("\n\nr: %b\ng: %b\nb: %b\n", r, g, b)
	fmt.Printf("\n\nDecimals\nr: %d\ng: %d\nb: %d\n", r, g, b)

	// Brighten color by 50% manually (Should bring it back to 30,155,255)

	r = r * 2
	g = g * 2
	b = b * 2

	color = (r << 16) + (g << 8) + b

	fmt.Println("\nLightened by 50%")
	fmt.Printf("\nColor: %b", color)
	fmt.Printf("\n\nr: %b\ng: %b\nb: %b\n", r, g, b)
	fmt.Printf("\n\nDecimals\nr: %d\ng: %d\nb: %d\n", r, g, b)
}
