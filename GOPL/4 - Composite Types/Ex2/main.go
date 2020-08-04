package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var mode = flag.Int("mode", 256, "Choose mode in list: sha256, sha384, sha512")

func main() {
	flag.Parse()
	var input string
	fmt.Scanf("%s", &input)
	content := []byte(input)
	switch *mode {
	case 256:
		fmt.Printf("SHA256: %x\n", sha256.Sum256(content))
		break
	case 384:
		fmt.Printf("SHA384: %x\n", sha512.Sum384(content))
		break
	case 512:
		fmt.Printf("SHA512: %x\n", sha512.Sum512(content))
		break
	default:
		fmt.Fprintf(os.Stderr, "Invalid mode sha%d\n", *mode)
		break
	}
}
