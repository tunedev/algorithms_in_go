package trie

import (
	"fmt"
	"strings"
)

// Trie - comes from ReTRIEval, it is another form of tree, other name for trie are Prefix, Digital, Radix ( tree)
// it is mostly used to implement auto completion
// trie allows us to store thousands even millions of words in minimal space and gives us optimal lookups
// Time complexity for opertions in a trie:
// Lookups: O(L) where L is the length of the word we are looking for
// Insert: O(L) where L is the length of the word we are inserting for
// Remove: O(L) where L is the length of the word we are looking for

const ALPHABET_SIZE = 26
const RUNE_A = 'a'

type node struct {
	val rune
	children map[rune]*node
	isEndOfWord bool
	actualWord string
}

func (n *node) String() string {
	return fmt.Sprintf("Value= %v\n", n.val)
}

func (n *node) HasChild(ch rune) bool {
	_, hasChild := n.children[ch - RUNE_A]
	return hasChild
}

func (n *node) AddChild(ch rune) {
	n.children[ch - RUNE_A] = &node{val: ch, children: make(map[rune]*node)}
}

func (n *node) GetChild(ch rune) *node {
	return n.children[ch - RUNE_A]
}

func (n *node) GetChildren() []*node {
	children := []*node{}
	for _, child := range n.children {
		children = append(children, child)
	}
	return children
}

type trie struct {
	root *node
}

func NewTrie() *trie {
	rootNode := &node{
		val: rune(' '),
		children: make(map[rune]*node),
	}
	return &trie{
		root: rootNode,
	}
}

func (t *trie) Insert (word string) {
	current := t.root
	sanitizedWord := t.sanitizeWord(word)
	for _, ch := range sanitizedWord {
		if !current.HasChild(ch){
			current.AddChild(ch)
		}
		current = current.GetChild(ch)
	}
	current.isEndOfWord = true
	current.actualWord = word
}

func (t *trie) sanitizeWord(word string) string {
	return strings.ToLower(strings.Replace(word, " ", "", -1))
}

func (t *trie) Contains(word string) bool {
	sanitizedWord := t.sanitizeWord(word)
	current := t.root
	for _, ch := range sanitizedWord {
		if !current.HasChild(ch){
			return false
		}
		current = current.GetChild(ch)
	}
	return current.isEndOfWord && current.actualWord == word
}

func (t *trie) FindSimilarWords(word string) []string {
	similarWords := []string{}
	lastNode := t.findLastNodeOf(word);
	t.findSimilarWords(lastNode, word, &similarWords)
	return similarWords
}

func (t *trie) Traverse() {
	root := t.root
	t.traverse(root)
}

func (t *trie) traverse(root *node) {
	if root == nil {
		return
	}
	fmt.Println(string(root.val))
	for _, child := range root.GetChildren(){
		t.traverse(child)
	}
}

func (t *trie) findSimilarWords(root *node, prefix string, existingWords *[]string) {
	if root == nil {
		return
	}
	if root.isEndOfWord {
		*existingWords = append(*existingWords, root.actualWord)
	}
	for _, child := range root.GetChildren() {
		t.findSimilarWords(child, prefix + string(child.val), existingWords)
	}
}

func (t *trie) findLastNodeOf(prefix string) *node {
	cur := t.root

	for _, ch := range prefix {
		child := cur.GetChild(ch);
		if child == nil {
			return nil
		}
		cur = child
	}
	return cur
}
