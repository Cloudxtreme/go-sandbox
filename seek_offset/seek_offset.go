package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

func ReadLastByte(r io.ReadSeeker) (b byte, err error) {
	if _, err = r.Seek(-1, os.SEEK_END); err != nil {
		return
	}

	if err = binary.Read(r, binary.LittleEndian, &b); err != nil {
		return
	}
	return
}

func main() {
	f, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ReadLastByte(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(b)
}
