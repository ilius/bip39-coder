package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

var wordCount = 2048
var bigRadix = big.NewInt(int64(wordCount))
var bigZero = big.NewInt(0)

var words, indexByWord = loadWords("english.txt")

func loadWords(fpath string) ([]string, map[string]int16) {
	textB, err := Asset(fpath)
	if err != nil {
		panic(err)
	}
	indexByWord := map[string]int16{}
	words := []string{}
	for wordIndex, wordB := range bytes.Split(textB, []byte("\n")) {
		word := string(wordB)
		if word == "" {
			continue
		}
		wordIndex16 := int16(wordIndex)
		if int(wordIndex16) != wordIndex {
			panic("Too many words")
		}
		indexByWord[word] = wordIndex16
		words = append(words, word)
	}
	return words, indexByWord
}

func decode(input string) []byte {
	answer := big.NewInt(0)
	j := big.NewInt(1)
	scratch := new(big.Int)

	for _, word := range strings.Split(input, " ") {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}
		wIndex, ok := indexByWord[word]
		if !ok {
			panic(fmt.Sprintf("Word %#v not found", word))
		}
		scratch.SetInt64(int64(wIndex))
		scratch.Mul(j, scratch)
		answer.Add(answer, scratch)
		j.Mul(j, bigRadix)
	}
	return answer.Bytes()
}

func encode(input []byte) string {
	x := new(big.Int)
	x.SetBytes(input)

	answer := []string{}
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, words[int16(mod.Int64())])
	}

	// leading zero bytes
	for _, i := range input {
		if i != 0 {
			break
		}
		answer = append(answer, words[0])
	}

	return strings.Join(answer, " ")
}

func main() {
	if len(words) != wordCount {
		panic("Mismatch number of words")
	}
	inputBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 1 && os.Args[1] == "-d" {
		os.Stdout.Write([]byte(encode(inputBytes) + "\n"))
	} else {
		os.Stdout.Write(decode(strings.Replace(strings.TrimSpace(string(inputBytes)), "\n", " ", -1)))
	}
}