// read_lbl Reads file line by line
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var fd *os.File
	var err error

	if len(os.Args) < 2 {
		fd = os.Stdin
	} else {
		fd, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()
	}

	s := bufio.NewScanner(fd)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
