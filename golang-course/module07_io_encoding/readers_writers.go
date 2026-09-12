package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// ─────── 1. io.Reader & io.Writer ───────
// These are arguably the two most important interfaces in the entire Go standard library.
//
// type Reader interface {
//     Read(p []byte) (n int, err error)
// }
//
// type Writer interface {
//     Write(p []byte) (n int, err error)
// }
//
// By programming against these interfaces instead of concrete types (like os.File or net.TCPConn),
// your code becomes infinitely reusable. A function that takes an `io.Reader` can read from
// a file, a network socket, an in-memory buffer, or an HTTP request body seamlessly.

func demonstrateReadersWriters() {
	fmt.Println("\n--- io.Reader & io.Writer ---")

	// 1. Reading from a string using strings.Reader
	// strings.Reader implements io.Reader
	r := strings.NewReader("Hello, Reader!")
	
	buf := make([]byte, 8)
	for {
		n, err := r.Read(buf)
		fmt.Printf("Read %d bytes: %q\n", n, buf[:n])
		
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading:", err)
			break
		}
	}

	// 2. Writing to an in-memory buffer using bytes.Buffer
	// bytes.Buffer implements both io.Reader and io.Writer
	var b bytes.Buffer
	
	// We can write strings or bytes to it
	b.Write([]byte("Hello, "))
	b.WriteString("Writer!\n")
	
	// fmt.Fprint families take an io.Writer!
	// (fmt.Print uses os.Stdout, which is an *os.File, which implements io.Writer)
	fmt.Fprintf(&b, "Formatted data: %d\n", 42)
	
	fmt.Print("Buffer contents: ", b.String())

	// 3. io.Copy
	// The ultimate Swiss Army knife. It copies from an io.Reader to an io.Writer
	// until EOF. It handles chunking automatically.
	r2 := strings.NewReader("Data to copy using io.Copy\n")
	var b2 bytes.Buffer
	
	copied, _ := io.Copy(&b2, r2)
	fmt.Printf("Copied %d bytes. Result: %s", copied, b2.String())
}
