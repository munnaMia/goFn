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

func InvertNumberTringleCol(n int) {
	for i := n; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

// Output
// 12345
// 1234
// 123
// 12
// 1

func RightTringle(n int, ch rune) {
	if n%2 == 0 {
		n++
	}

	for i := 1; i <= n; {
		for k := 1; k <= (n-i)/2; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
		i += 2
	}
}

// Output
//   *
//  ***
// *****

func InvertRightTringle(n int, ch rune) {
	if n%2 == 0 {
		n++
	}

	for i := n; i >= 1; {
		for k := 1; k <= (n-i)/2; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
		i -= 2
	}
}

// Output
//     *******
//      *****
//       ***
//        *

func SimpleDimond(n int, ch rune) {
	if n%2 == 0 {
		n++
	}

	RightTringle(n, ch)
	InvertRightTringle(n, ch)
}

// Output
//        *
//       ***
//      *****
//     *******
//     *******
//      *****
//       ***
//        *

func SideWallTringle(n int, ch rune) {
	Tringle(n, ch)
	InvertTringle(n-1, ch)
}

// Output
// *
// **
// ***
// **
// *
