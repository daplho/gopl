package main

import (
	"fmt"
	"sort"
)

type pslice []string

func (x pslice) Len() int           { return len(x) }
func (x pslice) Less(i, j int) bool { return x[i] < x[j] }
func (x pslice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func IsPalindrome(s sort.Interface) bool {
	isPalindrome := false
	for i, j := 0, s.Len()-1; i > j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		}
		isPalindrome = true
		break
	}

	return isPalindrome
}

func main() {
	var words = []string{"roots", "smash", "smash", "roots"}
	pwords := pslice(words)
	// 	sort.Sort(pwords)
	// 	fmt.Printf("%v\n", pwords)
	//
	// 	var reversepwords = make(pslice, len(pwords))
	// 	copy(reversepwords, pwords)
	// 	sort.Sort(sort.Reverse(reversepwords))
	// 	fmt.Printf("%v\n", reversepwords)
	if IsPalindrome(pwords) {
		fmt.Printf("'%v' is a palindrome", pwords)
	} else {
		fmt.Printf("'%v' is not a palindrome", pwords)
	}
}
