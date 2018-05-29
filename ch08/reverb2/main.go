// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup // number of active echo goroutines
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go echo(c, wg, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()

	go func() {
		wg.Wait()
		c.(*net.TCPConn).CloseWrite()
	}()
}

func echo(c net.Conn, wg sync.WaitGroup, shout string, delay time.Duration) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
