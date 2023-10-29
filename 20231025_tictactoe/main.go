package main

import (
	"errors"
	"fmt"
)

func main() {

	fields := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 2}}

	n := 3

	winner, err := isWinner(n, fields)

	if err == nil {

		fmt.Println("Did we win?", func(winner bool) string {
			if winner {
				return "yes!"
			} else {
				return "no..."
			}
		}(winner))
	} else {
		fmt.Println("You provided invalid input for a game board of", n, "x", n, "elements! Error:", err.Error()+"!")
	}

}

func isWinner(n int, fields [][]int) (bool, error) {
	rows := make([]int, n)
	cols := make([]int, n)
	diag := 0
	adiag := 0

	for _, v := range fields {
		r, c := v[0], v[1]
		if r > n-1 || c > n-1 {
			return false, errors.New("coordinates are outside of game board")
		}

		rows[r]++
		cols[c]++

		if r == c {
			diag++
		}

		if r+c == n-1 {
			adiag++
		}

	}

	// fmt.Println("rows:", rows)
	// fmt.Println("cols:", cols)
	// fmt.Println("diag:", diag)
	// fmt.Println("adiag:", adiag)

	if diag == n || adiag == n {
		return true, nil
	}

	for _, v := range rows {
		if v == n {
			return true, nil
		}
	}

	for _, v := range cols {
		if v == n {
			return true, nil
		}
	}

	return false, nil
}
