package invertedtriangle

import "fmt"

func solution(n int) string {
	triangle := ``

	if n%2 == 1 && n > 0 {
		height := (n / 2) + 1
		space := 0
		stars := n
		nNew := n

		for height > 0 {
			spaceChar := space
			for spaceChar > 0 {
				triangle += fmt.Sprintf(` `)
				spaceChar--
			}

			for stars > 0 {
				triangle += fmt.Sprintf(`*`)
				stars--
			}

			if nNew != n {
				spaceChar = space
				for spaceChar > 0 {
					triangle += fmt.Sprintf(` `)
					spaceChar--
				}
			}

			height--
			nNew -= 2
			stars = nNew
			space = (n - nNew) / 2
		}
	}

	return triangle
}
