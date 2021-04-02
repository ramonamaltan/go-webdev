package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

// hmac = keyed-hash message authentication code
// hash: same input produces the same output

func main() {
	c := getCode("test1@example.com")
	fmt.Println(c)
	c = getCode("test1@example.com")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
