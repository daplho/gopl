package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var x io.Writer
	x = os.Stdout
	fmt.Printf("os.Stdout's type is %T\n", x)

	// possibility 1: asserted type `T` is a concrete type
	// (`*os.File` and `*bytes.Buffer` respectively below)
	if f, ok := x.(*os.File); ok { // success: f == os.Stdout
		fmt.Println("Successfully asserted 'x' to an '*os.File'")
		fmt.Printf("os.Stdout's name is %s\n", f.Name())
	}

	if c, ok := x.(*bytes.Buffer); !ok {
		fmt.Println("panic: interface holds *os.File, not *bytes.Buffer (this is expected)")
	} else {
		fmt.Printf("os.Stdout asserted to a '*bytes.Buffer' contains: %s", c.String())
	}

	// possibility 2: asserted type `T` is an interface type
	//	var y io.Writer
	//	y = os.Stdout
	//	rw := y.(io.ReadWriter) // success: *os.File has both Read and Write
	//	fmt.Printf("os.Stdout's type is %T\n", rw)
	//
	//	y = new(ByteCounter)
	//	rw = y.(io.ReadWriter) // panic: *ByteCounter has no Read method
}
