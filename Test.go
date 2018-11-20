package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := bytes.IndexByte([]byte("Golang is Good"),byte('G'))
	fmt.Println("Result: s",s)
}
