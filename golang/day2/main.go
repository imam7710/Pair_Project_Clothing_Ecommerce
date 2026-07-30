package main

import (
	"fmt"
)

func main() {
	balances := [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	// array teluar ada 2 baris dan array terdalam ada 5 kolom
	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}
