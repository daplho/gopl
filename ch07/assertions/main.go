package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {
	var x io.Writer
	x = os.Stdout
	fmt.Printf("os.Stdout's type is %T\n", x)

	// possibility 1: asserted type `T` is a concrete type
	// (`*os.File` and `*bytes.Buffer` respectively below)
	if f, ok := x.(*os.File); ok {
		fmt.Println("success: f == os.Stdout")
		fmt.Printf("os.Stdout's name is %s\n", f.Name())
	}

	if c, ok := x.(*bytes.Buffer); !ok {
		fmt.Println("panic: interface holds *os.File, not *bytes.Buffer (this is expected)")
	} else {
		fmt.Printf("os.Stdout asserted to a '*bytes.Buffer' contains: %s", c.String())
	}

	// possibility 2: asserted type `T` is an interface type
	var y io.Writer
	y = os.Stdout

	if rw, ok := y.(io.ReadWriter); ok {
		fmt.Println("success: *os.File has both Read and Write (this is expected)")
		rw.Write([]byte("Hello Assertions!\n"))
	}

	y = new(ByteCounter)
	if rw, ok := y.(io.ReadWriter); !ok {
		fmt.Println("panic: *ByteCounter has no Read method (this is expected)")
	} else {
		rw.Write([]byte("Goodbye Assertions!\n"))
	}
}
