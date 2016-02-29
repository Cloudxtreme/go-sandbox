package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

var (
	tc     = flag.Bool("t", false, "Show truecolor only")
	size   = flag.Int("s", 51, "Size of truecolor cube [26 - 51]")
	pcolor = flag.Bool("p", false, "Print RGB color code")
	code   = flag.Bool("c", false, "With 256 code")
	list   = flag.Bool("l", false, "Print list style")
	str    = flag.String("str", "", "Print string")
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
	flag.Parse()

	if !*tc {

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
						if *pcolor {
							CsiPrintf(fgbg, color, fmt.Sprintf("%03s", strconv.Itoa(color)))
						} else {
							CsiPrintf(fgbg, color, "::")
						}
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

	}

	switch {
	case *size < 26:
		log.Fatal("cube size should be more then 26")
	case *size > 91:
		log.Fatal("cube size should be less then 51")
	default:

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

			var r, g, b, gi, ri, bi int
			for g, gi = 0, 0; g < 256; g += *size {
				gi++
				for r, ri = 0, 0; r < 256; r += *size {
					ri++
					for b, bi = 0, 0; b < 256; b += *size {
						bi++
						if *pcolor {
							color := 16 + ri*36 + gi*6 + bi - 43
							fmt.Printf("%d :", color)
							CsiPrintfTrueColor(fgbg, r, g, b, fmt.Sprintf(" S: %03s %03s %03s :E ", strconv.Itoa(r), strconv.Itoa(g), strconv.Itoa(b)))
							fmt.Printf("\n")
						} else if *str != "" {
							CsiPrintfTrueColor(fgbg, r, g, b, *str)
							fmt.Printf("\n")
						} else {
							CsiPrintfTrueColor(fgbg, r, g, b, "::")
						}
						if *list {
							fmt.Printf("\n")
						}
					}
					// fmt.Printf(" ")
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
}
