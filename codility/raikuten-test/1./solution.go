package lesson

// import "fmt"
import "strconv"

// Solution 1
func Solution(S string) string {
	// write your code in Go 1.4
	formattedS:=""
	count:=0
	lenS:=len(S)

	for i, v := range S {

		if v != 32 && v!= 45 {	
			formattedS += strconv.Itoa(int(v) - 48)
			count++	

			if count == 3 && i < lenS-1 {
				formattedS += "-"
				count = 0
			}
		} 
		
	}

	if count == 1 {
		lenFS := len(formattedS)
		r := []rune(formattedS)

		r[lenFS-2], r[lenFS-3] = r[lenFS-3], r[lenFS-2]

		return string(r)
	}

	return formattedS
	
}
