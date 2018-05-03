package main

import (
	_ "bytes"
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
	f := x.(*os.File) // success: f == os.Stdout
	fmt.Printf("os.Stdout's name is %s\n", f.Name())
	// c := x.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	// fmt.Printf("os.Stdout's contents are %s\n", c.String())

	// possibility 2: asserted type `T` is an interface type
}
