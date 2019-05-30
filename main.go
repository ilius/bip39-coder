package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/ilius/bip39-coder/bip39"
)

func main() {
	inputBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 1 && os.Args[1] == "-d" {
		os.Stdout.Write([]byte(bip39.Encode(inputBytes) + "\n"))
	} else {
		os.Stdout.Write(bip39.Decode(strings.Replace(strings.TrimSpace(string(inputBytes)), "\n", " ", -1)))
	}
}
