package main

import "fmt"

const NMAX int = 5

type table [NMAX]rune

func isiArray(t *table, n *int) {
	var berhenti bool = false
	*n = 0

	for !berhenti && *n < NMAX {
		fmt.Scanf("%c", &t[*n])
		fmt.Println(t[*n])

		if t[*n] == 'b' {
			fmt.Println("TRIGGER")
			berhenti = true
		}
		*n++
	}
}

func main() {
	var tab table
	var n int

	isiArray(&tab, &n)

	for i := 0; i < n; i++ {
		fmt.Print(string(tab[i])) // Print as string to display characters
	}
}
