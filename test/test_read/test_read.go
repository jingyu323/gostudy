package main

import "fmt"

func solution(n int, pages [7]int) {

	var result int

	// TODO: 请在此编写代码
	for n > 0 {

		for i := 0; i < len(pages); i++ {
			result = i + 1
			n = n - pages[i]
			if n <= 0 {
				break
			}

		}

	}
	fmt.Print(result)

}

func main() {
	var n int
	fmt.Scan(&n)
	var pages [7]int

	for i := 0; i < 7; i++ {
		fmt.Scan(&pages[i])
	}

	solution(n, pages)
}
