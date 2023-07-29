package main

import (
	"fmt"

	t "github.com/tunedev/algorithms_in_go/trie"
)

func main() {
	trie := t.NewTrie()
	trie.Insert("cat")
	trie.Insert("canada")
	trie.Insert("cat-stuffs")
	trie.Insert("Cat and Stuffs")
	trie.Insert("Cotainers")
	trie.Insert("Cotainers and Shipments")

	fmt.Println(trie.FindSimilarWords("co"))
}
