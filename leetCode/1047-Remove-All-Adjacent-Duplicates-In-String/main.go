package main

import "fmt"

func main() {
	example1 := "abbaca"
	example2 := "azxxzy"
	example3 := "aaaaaaaaa"

	res1 := removeDuplicates(example1)
	res2 := removeDuplicates(example2)
	res3 := removeDuplicates(example3)

	fmt.Printf("res 1: [%s]\n\n", res1)
	fmt.Printf("res 2: [%s]\n\n", res2)
	fmt.Printf("res 3: [%s]\n\n", res3)
}

func removeDuplicates(input string) string {

	buf := make([]byte, len(input))

	checkPoint := 0
	recordPoint := 0

	for i := 0; i < len(input); i++ {

		if buf[0] == 0 && i < len(input)-1 {
			buf[0] = input[i]
			recordPoint++
			i++
		}

		if input[i] != buf[checkPoint] {
			buf[recordPoint] = input[i]
			checkPoint++
			recordPoint++

		} else {
			buf[checkPoint] = 0

			if checkPoint != 0 {
				checkPoint--
				recordPoint--
			} else {
				recordPoint = 0
			}

		}
	}

	buf = buf[:recordPoint]

	return string(buf)
}
