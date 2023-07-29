package trie

// import (
// 	"fmt"
// 	"strings"
// )

// // Trie - comes from ReTRIEval, it is another form of tree, other name for trie are Prefix, Digital, Radix ( tree)
// // it is mostly used to implement auto completion
// // trie allows us to store thousands even millions of words in minimal space and gives us optimal lookups
// // Time complexity for opertions in a trie:
// // Lookups: O(L) where L is the length of the word we are looking for
// // Insert: O(L) where L is the length of the word we are inserting for
// // Remove: O(L) where L is the length of the word we are looking for

// const ALPHABET_SIZE = 26
// const RUNE_A = 'a'

// type node struct {
// 	val rune
// 	children [ALPHABET_SIZE]*node
// 	isEndOfWord bool
// 	actualWord string
// }

// func (n *node) String() string {
// 	return fmt.Sprintf("Value= %v\n", n.val)
// }

// func (n *node) HasChild(ch rune) bool {
// 	return n.children[ch - RUNE_A] != nil
// }

// func (n *node) AddChild(ch rune) {
// 	n.children[ch - RUNE_A] = &node{val: ch}
// }

// func (n *node) GetChild(ch rune) *node {
// 	return n.children[ch - RUNE_A]
// }

// type trie struct {
// 	root *node
// }

// func NewTrie() *trie {
// 	rootNode := &node{
// 		val: rune(' '),
// 	}
// 	return &trie{
// 		root: rootNode,
// 	}
// }

// func (t *trie) Insert (word string) {
// 	current := t.root
// 	sanitizedWord := t.sanitizeWord(word)
// 	for _, ch := range sanitizedWord {
// 		if !current.HasChild(ch){
// 			current.AddChild(ch)
// 		}
// 		current = current.GetChild(ch)
// 	}
// 	current.isEndOfWord = true
// 	current.actualWord = word
// }

// func (t *trie) sanitizeWord(word string) string {
// 	return strings.ToLower(strings.Replace(word, " ", "", -1))
// }

// func (t *trie) Contains(word string) bool {
// 	sanitizedWord := t.sanitizeWord(word)
// 	current := t.root
// 	for _, ch := range sanitizedWord {
// 		if !current.HasChild(ch){
// 			return false
// 		}
// 		current = current.GetChild(ch)
// 	}
// 	return current.isEndOfWord && current.actualWord == word
// }
