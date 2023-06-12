package main

import (
	"fmt"
	"strings"

	stack "github.com/tunedev/algorithms_in_go/stack"
)

func main() {
	fmt.Println(reverseString("Hallo"))
	fmt.Println(isExpressionBalanced("(((()))"))
}

func reverseString(val string) string{
	s1 := stack.NewStack[rune](len(val))
	var result string
	for _, s := range val{
		s1.Push(s)
	}
	for !s1.IsEmpty() {
		result += string(s1.Pop())
	}
	return result
}

func isExpressionBalanced(exp string) bool {
	openingBrackets := "([{<"
	closingBrackets := ")]}>"

	isOpenedBracket := func(cha string) bool {
		return strings.Contains(openingBrackets, cha)
	}
	isClosedBracket := func(cha string) bool {
		return strings.Contains(closingBrackets, cha)
	}

	s1 := stack.NewStack[rune](len(exp))
	for _, s := range exp {
		if isOpenedBracket(string(s)){
			s1.Push(s)
		}
		if isClosedBracket(string(s)){
			if strings.IndexRune(closingBrackets, s) != strings.IndexRune(openingBrackets, s1.Pop()) {
				return false
			}
		}
	}
	return s1.IsEmpty()
}