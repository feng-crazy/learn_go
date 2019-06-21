package main

import (
	"bytes"
	"fmt"
	"encoding/gob"
	"log"
	"os"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer   // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network)	// Will read from network.
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	fmt.Println("network:", network)
	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("q :%s: {%d,%d}\n", q.Name, *q.X, *q.Y)

	file, _ := os.OpenFile("test.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	fenc := gob.NewEncoder(file)
	ferr := fenc.Encode(P{3, 4, 5, "Pythagoras"})
	if ferr != nil {
		log.Println("Error in encoding gob")
	}
}