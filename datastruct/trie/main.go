// impl. from Tim McNamara (timclicks) https://www.youtube.com/watch?v=f9B87LA86g0

package main

import (
	. "exercises/shared"
	"fmt"

	"github.com/ajitid/litter"
)

type node struct {
	atEnd    bool
	children map[rune]*node
}

func newNode() *node {
	return &node{
		children: make(map[rune]*node),
	}
}

type trie struct {
	root *node
}

func newTrie() trie {
	t := trie{}
	t.root = newNode()
	return t
}

func (t *trie) insert(text string) {
	currentNode := t.root

	for _, char := range text {
		if _, ok := currentNode.children[char]; !ok {
			currentNode.children[char] = newNode()
		}
		currentNode = currentNode.children[char]
	}

	currentNode.atEnd = true
}

func (t trie) contains(text string) bool {
	currentNode := t.root

	for _, char := range text {
		node, ok := currentNode.children[char]
		if !ok {
			return false
		}
		currentNode = node
	}

	return currentNode.atEnd
}

func main() {
	PreMain()

	words := newTrie()
	words.insert("hi")
	words.insert("hey")

	fmt.Println("hey", words.contains("hey"))
	fmt.Println("heyy", words.contains("heyy"))
	fmt.Println("he", words.contains("he"))

	fmt.Println()
	litter.Dump(words)
}
