package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Trie struct {
	val  byte
	sons [200]*Trie
	end  int
}

func (t *Trie) Insert(word string) {
	node := t
	size := len(word)
	for i := 0; i < size; i++ {
		idx := int(word[i])
		if node.sons[idx] == nil {
			node.sons[idx] = &Trie{val: word[i]}
		}

		node = node.sons[idx]
	}

	node.end++
}

func (t *Trie) HasDirtyWord(input string) bool {
	node := t
	size := len(input)
	for start := 0; start < size; start++ {
		// dirty word candidate is input[start:size]
		node = t
		for i := start; i < size; i++ {
			idx := int(input[i])
			if node.sons[idx] == nil {
				break
			}

			node = node.sons[idx]
			if node.end > 0 {
				return true // match dirty word
			}
		}
	}

	return false
}

func readLine(reader *bufio.Reader) string {
	res, _ := reader.ReadString('\n')
	res = strings.TrimSpace(res)
	return res
}

func main() {
	trie := Trie{}

	reader := bufio.NewReader(os.Stdin)

	var n, m int
	var dirtyWord, testLine string
	var tmpRes string
	tmpRes = readLine(reader)
	n, _ = strconv.Atoi(tmpRes)

	for i := 0; i < n; i++ {
		dirtyWord = readLine(reader)
		trie.Insert(dirtyWord)
	}

	tmpRes = readLine(reader)
	m, _ = strconv.Atoi(tmpRes)
	for i := 0; i < m; i++ {
		testLine = readLine(reader)
		if trie.HasDirtyWord(testLine) {
			fmt.Println("True")
			continue
		}
		fmt.Println("False")
	}
}
