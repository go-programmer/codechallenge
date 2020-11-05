package lesson

// Solution 2
func Solution(T []int) string {
	// write your code in Go 1.4

	lenT:=len(T)
	temperatures:= lenT/4
	ha:= make([]int, 4)
	season := 0

	innerCounter := 1

	for i := 0; i < lenT; i=i+temperatures {
		h:= T[i]
		l:= T[i]

		for j := 1; j < temperatures; j++ {
			if T[innerCounter] > h {
				h = T[innerCounter]
			}

			if T[innerCounter] < l{
				l = T[innerCounter]
			}
			innerCounter++
		}

		ha[season] = h-l
		season++
		innerCounter++

	}


	seasonM:= make(map[int]string)
	seasonM[0] = "WINTER"
	seasonM[1] = "SPRING"
	seasonM[2] = "SUMMER"
	seasonM[3] = "AUTUMN"

	season=0
	h:=ha[0]

	for i:=1; i < 4 ; i++ {
		
		if ha[i] > h {
			season = i
		}
	}

	
	return seasonM[season]
	
}
