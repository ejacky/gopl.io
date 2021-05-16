package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

var hashMethod = flag.String("hash-method", "sha256", "the hash method, support sha256, sha384, sha512 and default is sha256")

// echo "x" | go run main.go

func main() {
	flag.Parse()

	var raw string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		raw += input.Text()
	}

	switch strings.ToLower(*hashMethod) {
	case "sha512":
		fmt.Printf("%x", sha512.Sum512([]byte(raw)))
	case "384":
		fmt.Println("sorry, don't have a method to implement it")
	default:
		fmt.Printf("%x", sha256.Sum256([]byte(raw)))

	}

}
