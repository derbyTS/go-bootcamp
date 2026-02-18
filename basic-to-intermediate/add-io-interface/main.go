package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// -------------------------------------------------------------------------
// 1. THE CONTRACT (The Interface)
// In Java: interface ReaderInterface { int RRead(byte[] p); }
// In Go:   No "implements" keyword is needed. If a struct has this method,
//
//	it ALREADY is a ReaderInterface. This is "Static Duck Typing."
//
// -------------------------------------------------------------------------
type ReaderInterface interface {
	RRead(p []byte) (n int, err error)
}

// -------------------------------------------------------------------------
// 2. THE DATA (The Struct)
// This is just a "dumb" box of data. It doesn't know about interfaces yet.
// -------------------------------------------------------------------------
type RReader struct {
	s        string
	i        int64 // tracks where we are in the string
	prevRune int   // tracks last read position
}

// -------------------------------------------------------------------------
// 3. THE CONSTRUCTOR (The Helper)
// We return a *Pointer (&) because:
// - It's efficient (don't copy the whole string/struct).
// - It allows the RRead method to update the index 'i' on the ORIGINAL struct.
// -------------------------------------------------------------------------
func NNewReader(s string) *RReader {
	return &RReader{s: s, i: 0, prevRune: -1}
}

// -------------------------------------------------------------------------
// 4. THE BEHAVIOR (The Method)
// (r *RReader) is the "Receiver." It's like 'this' in Java.
// Because this method signature matches ReaderInterface exactly,
// *RReader now "implements" ReaderInterface.
// -------------------------------------------------------------------------
func (r *RReader) RRead(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF // Stop if we've read everything
	}
	r.prevRune = -1

	// Copy data from our internal string (r.s) into the provided buffer (b)
	n = copy(b, r.s[r.i:])

	// IMPORTANT: We update 'i' on the pointer 'r'.
	// If 'r' wasn't a pointer, 'i' would reset every time!
	r.i += int64(n)
	return
}

// -------------------------------------------------------------------------
// 5. THE CONSUMER (The Function)
// This function accepts ANY type that has a RRead method.
// It doesn't care if it's a struct, a pointer, or a custom string type.
// -------------------------------------------------------------------------
func readFromReader(r ReaderInterface) {
	buf := make([]byte, 1024)
	n, err := r.RRead(buf) // Polymorphism in action
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Result:", string(buf[:n]))
}

// This is where I get the idea to implement it all manually - strings.NewReader is function but it implements the Read
func readFromReaderFromTutorial(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Result:", string(buf[:n]))
}

func main() {
	// NNewReader returns a *RReader (pointer).
	// Since *RReader has the RRead method, it satisfies ReaderInterface.
	readFromReader(NNewReader("Hello from the Pointer!"))

	readFromReaderFromTutorial(strings.NewReader("Hello Reader"))
}
