package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// nopqrstuvwxyzabcdefghijklm
var m = map[byte]byte{
	'n': 'a',
	'o': 'b',
	'p': 'c',
	'q': 'd',
	'r': 'e',
	's': 'f',
	't': 'g',
	'u': 'h',
	'v': 'i',
	'w': 'j',
	'x': 'k',
	'y': 'l',
	'z': 'm',
	'a': 'n',
	'b': 'o',
	'c': 'p',
	'd': 'q',
	'e': 'r',
	'f': 's',
	'g': 't',
	'h': 'u',
	'i': 'v',
	'j': 'w',
	'k': 'x',
	'l': 'y',
	'm': 'z',
}

func (reader rot13Reader) Read(bytes []byte) (int, error) {
	size, err := reader.r.Read(bytes)
	//fmt.Printf("Before: %v, %v\n", size, bytes[:size])
	for i, b := range bytes {
		newval := m[b]
		if newval != 0 {
			bytes[i] = m[b]
		}
	}
	//fmt.Printf("After: %v, %v\n", size, bytes[:size])
	return size, err
}

func main() {
	s := strings.NewReader("lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
