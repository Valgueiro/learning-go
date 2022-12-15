package main

import "fmt"

func main() {
	fmt.Print(isValid("({)}"))
}

func isOpeningBrackets(b rune) bool {
	return b == '(' || b == '[' || b == '{'
}

func isClosingCorrectly(lastOpeningBracket rune, actualB rune) bool {
	rightClosing := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	right, ok := rightClosing[lastOpeningBracket]

	if ok {
		return actualB == right
	}

	return false
}

func isValid(s string) bool {
	stack := Stack{}
	for _, bracket := range s {
		if isOpeningBrackets(bracket) {
			stack.push(bracket)
		} else if !isClosingCorrectly(stack.pop(), bracket) {
			return false
		}
	}
	return stack.isEmpty()
}

type Stack struct {
	Elements []rune
}

func (s *Stack) isEmpty() bool {
	return len(s.Elements) == 0
}

func (s *Stack) push(elem rune) {
	s.Elements = append(s.Elements, elem)
}

func (s *Stack) pop() rune {
	if s.isEmpty() {
		return -1
	}
	out := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return out
}

func (s *Stack) print() {
	fmt.Println(s.Elements)
}
