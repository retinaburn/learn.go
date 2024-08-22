package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (myReader MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	//copy(b, "A")
	return 1, nil
}

func main() {
	var mybytes [200]byte
	r := MyReader{}
	r.Read(mybytes[:])

	fmt.Printf("cap = %v, %v", cap(mybytes), mybytes)
	reader.Validate(MyReader{})
}
