package main

import "fmt"

func reverseUsingString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func reverseUsiingSlices(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func main() {
	str := "anit"
	strRev := reverseUsiingSlices(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
