package main

import "fmt"

func decreasePattern(n int) {
	for i := n; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func trianglePattern(n int) {
	//n=5
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func isLuhan(nums int) bool {
	num := doubleNumbers(nums)
	fmt.Println("Sum:", num)
	return num/10 == 0
}

func doubleNumbers(n int) int {
	double := n * 2
	sum := 0
	if double > 10 {
		sum = 1 + n%10
	} else {
		sum = double
	}
	fmt.Println("doubled number:", double)
	return sum
}

func centeredHash(rows int) {
	// maxWidth := rows
	for i := rows; i > 1; i-- {
		if i%2 == 0 {
			space := (rows - i) / 2
			for s := 0; s < space; s++ {
				fmt.Print(" ")
			}
			for j := 0; j < i; j++ {
				fmt.Print("#")
			}

			fmt.Println()
		}
	}
}

func diamondShape(rows int) {
	for i := 1; i < rows; i++ {
		if i%2 == 0 {
			space := (rows - i) / 2
			for s := 0; s < space; s++ {
				fmt.Print(" ")
			}
			for j := 0; j < i; j++ {
				fmt.Print("#")
			}

			fmt.Println()
		}
	}
	for k := rows - 1; k > 1; k-- {
		if k%2 == 0 {
			space := (rows - k) / 2
			for s := 0; s < space; s++ {
				fmt.Print(" ")
			}
			for j := 0; j < k; j++ {
				fmt.Print("#")
			}
			fmt.Println()
		}
	}
}

func crusedDimoamen(rows int) {
	for i := 0; i <= rows; i++ {
		for s := 0; s <= i; s++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}

		fmt.Print()
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		
		fmt.Println()
	}
	fmt.Print()
	for i := 1; i <= rows; i++ {
		for s := 0; s < rows-i; s++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	
}
func main() {
	// decreasePattern(5)
	// trianglePattern(4)
	// fmt.Println(doubleNumbers(349))
	// fmt.Println(isLuhan(349))
	// centeredHash(8)
	// diamondShape(10)
	crusedDimoamen(3)
}
