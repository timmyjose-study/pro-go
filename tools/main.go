package main

import "fmt"

func main() {
	PrintHello()
	for i := 1; i <= 5; i++ {
		PrintNumber(i)
	}
}

// revive:disable:exported
func PrintHello() {
	fmt.Println("Hello, Go!")
}

// revive:enable:exported
// PrintNumber prints, and this may be shocking, a number.
func PrintNumber(n int) {
	fmt.Println(n)
}
