package main

import "fmt"

func main() {
	words := "selamat malam"

	res := map[string]int{}

	for _, l := range words {
		letter := string(l)

		value, exist := res[letter]
		if exist {
			res[letter] = value + 1
		} else {
			res[letter] = 1
		}
		fmt.Println(letter)
	}
	fmt.Println(res)
}
