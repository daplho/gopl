package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*c += WordCounter(count)

	return count, nil
}

func main() {
	var c WordCounter
	c.Write([]byte("Spicy jalapeno pastrami ut ham turducken.\n Lorem sed ullamco, leberkas sint short loin strip steak ut shoulder shankle porchetta venison prosciutto turducken swine.\n"))
	fmt.Println(c)

	c = 0
	var words = "Deserunt kevin frankfurter tongue aliqua incididunt tri-tip shank nostrud.\n"
	fmt.Fprintf(&c, "%s", words)
	fmt.Println(c)
}
