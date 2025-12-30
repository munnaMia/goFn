package patterns

import "fmt"

func Square(n int, ch rune) {
	for range n {
		for range n {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}

// Output
// **********
// **********
// **********
// **********
// **********
// **********
// **********
// **********
// **********
// **********

func Tringle(n int, ch rune) {
	for i := range n {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}

// Output
// *
// **
// ***
// ****
// *****
// ******
// *******
// ********
// *********
// **********

func NumberTringleCol(n int) {
	for i := range n {
		for j := 0; j <= i; j++ {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
}

// Output
// 1
// 12
// 123
// 1234
// 12345

func NumberTringleRow(n int) {
	for i := range n {
		for j := 0; j <= i; j++ {
			fmt.Print(i + 1)
		}
		fmt.Println()
	}
}

// Output
// 1
// 22
// 333
// 4444
// 55555

func InvertTringle(n int, ch rune) {
	for i := range n {
		for j := n - 1; j >= i; j-- {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}

// Output
// *****
// ****
// ***
// **
// *
