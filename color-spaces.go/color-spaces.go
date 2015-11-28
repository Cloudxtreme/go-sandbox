package main

import (
	"fmt"
	"strconv"
)

// Standard color escape sequence printf
// spec: print "\x1b[${fgbg};5;${color}m${string}";
// e.g.: print "\x1b[     38;5;     166m${string}";
func CsiPrintf(fgbg int, color int, character string) {
	fmt.Printf("\x1b[%d;5;%dm%s\x1b[0m", fgbg, color, character)
}

func CsiPrintfReverse(fgbg int, color int, character string) string {
	switch fgbg {
	case 38:
		fgbg += 10
	case 48:
		fgbg -= 10
	}
	switch color > 244 {
	case true:
		color += 23
	case false:
		color -= 23
	}
	out := fmt.Sprintf("\x1b[%d;5;%dm%s\x1b[0m", fgbg, color, character)

	return out
}

// True color escape sequence printf
// spec: print "\x1b[${fgbg};2;${red};${green};${blue}m${string}\x1b[0m"
// e.g.: print "\x1b[     38;2;   255;     100;      0mTRUECOLOR\x1b[0m"
func CsiPrintfTrueColor(fgbg int, r int, g int, b int, character string) {
	fmt.Printf("\x1b[%d;2;%d;%d;%dm%s\x1b[0m", fgbg, r, g, b, character)
}

func lineSeparater(line int) {
	for i := 0; i < line; i++ {
		fmt.Printf("\n")
	}
}

func main() {
	fmt.Println("256 color mode\n")

	for fgbg := 38; fgbg <= 48; fgbg += 10 {
		switch fgbg {
		case 38:
			fmt.Println("Foreground Color cube\n")
		case 48:
			fmt.Println("Background Color cube\n")
		}

		fmt.Println("System colors:")
		for color := 0; color < 8; color++ {
			CsiPrintf(fgbg, color, "::")
		}
		lineSeparater(1)
		for color := 8; color < 16; color++ {
			CsiPrintf(fgbg, color, "::")
		}

		lineSeparater(2)

		var r, g, b int
		for g = 0; g < 6; g++ {
			for r = 0; r < 6; r++ {
				for b = 0; b < 6; b++ {
					color := 16 + r*36 + g*6 + b
					CsiPrintf(fgbg, color, "::")
				}
				fmt.Printf(" ")
			}
			lineSeparater(1)
		}
		lineSeparater(1)

		fmt.Println("Grayscale ramp:")
		for gray := 232; gray < 256; gray++ {
			CsiPrintf(fgbg, gray, strconv.Itoa(gray)+" ")
		}
		lineSeparater(2)
	}

	fmt.Println("24bit True color mode\n")

	for fgbg := 38; fgbg <= 48; fgbg += 10 {
		switch fgbg {
		case 38:
			fmt.Println("Foreground Color cube\n")
		case 48:
			fmt.Println("Background Color cube\n")
		}

		fmt.Println("System colors:")
		for color := 0; color < 8; color++ {
			CsiPrintf(fgbg, color, "::")
		}
		lineSeparater(1)
		for color := 8; color < 16; color++ {
			CsiPrintf(fgbg, color, "::")
		}

		lineSeparater(2)

		var r, g, b int
		for g = 0; g < 256; g += 51 {
			for r = 0; r < 256; r += 51 {
				for b = 0; b < 256; b += 51 {
					CsiPrintfTrueColor(fgbg, r, g, b, "::")
				}
				fmt.Printf(" ")
			}
			lineSeparater(1)
		}
		lineSeparater(1)

		fmt.Println("Grayscale ramp:")
		for gray := 8; gray < 256; gray += 10 {
			CsiPrintfTrueColor(fgbg, gray, gray, gray, strconv.Itoa(gray)+" ")
		}
		lineSeparater(2)
	}
}