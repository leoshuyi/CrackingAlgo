package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Trie struct {
	val  byte
	sons [256]*Trie
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
	res = strings.TrimSuffix(res, "\n")
	return res
}

func main() {
	trie := Trie{}

	file, _ := os.Open("8.in")
	defer file.Close()

	outfile, _ := os.OpenFile(
		"x8.out",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	reader := bufio.NewReader(file)

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
			//fmt.Println("True")
			outfile.Write([]byte("True\n"))
			continue
		}
		//fmt.Println("False")
		outfile.Write([]byte("False\n"))
	}
}
