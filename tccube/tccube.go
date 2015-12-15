// tccube Displays the 24bit color cube
package main

import (
	"flag"
	"fmt"
	"strconv"
)

var (
	r, g, b int
	only    = flag.String("only", "", "Display only the [fg, bg] cube -- Default is display the both")
	size    = flag.Int("s", 8, "Size of truecolor cube [26 ~ 51]")
	crange  = flag.Int("range", 256, "Display color range")
)

// Konsole based truecolor escape sequence printf
// spec: print "\x1b[${fgbg};2;${red};${green};${blue}m${character}\x1b[0m"
// e.g.: print "\x1b[     38;2;   255;     100;      0m TRUECOLOR  \x1b[0m"
func SgrPrintfTrueColor(fgbg int, r int, g int, b int, character string) {
	fmt.Printf("\x1b[%d;2;%d;%d;%dm%s\x1b[0m", fgbg, r, g, b, character)
}

func lineSeparater(line int) {
	for i := 0; i < line; i++ {
		fmt.Printf("\n")
	}
}

func main() {
	flag.Parse()

	fmt.Println("24bit True color mode")

	for fgbg := 38; fgbg <= 48; fgbg += 10 {

		// parse disable [fg,bg] cube
		if *only != "" && *only == "fg" && fgbg == 48 {
			continue
		} else if *only == "bg" && fgbg == 38 {
			continue
		}

		switch fgbg {
		case 38:
			fmt.Println("\nForeground Color cube\n")
		case 48:
			fmt.Println("\nBackground Color cube\n")
		}

		for i := 0; i < 256; i += 64 {
			for g = i; g < i+64; g += *size {
				for r = i; r < i+64; r += *size {
					for b = i; b < i+64; b += *size {
						SgrPrintfTrueColor(fgbg, r, g, b, fmt.Sprintf(" %03s %03s %03s ", strconv.Itoa(r), strconv.Itoa(g), strconv.Itoa(b)))
					}
					lineSeparater(1)
				}
			}
		}
	}
}
