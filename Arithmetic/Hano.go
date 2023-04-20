package Arithmetic

import "fmt"

// 经典的暴力递归思想题目
func Hano1(n int) {
	leftToRight(n)
}

func leftToRight(n int) {
	if n == 1 {
		fmt.Println("move 1 from left to right")
		return
	}
	leftToMid(n - 1)
	fmt.Println("move", n, "from left to right")
	midToRight(n - 1)
}
func leftToMid(n int) {
	if n == 1 {
		fmt.Println("move 1 from left to mid")
	}
	leftToRight(n - 1)
	fmt.Println("move", n, "from left to right")
	rightToMid(n - 1)
}
func rightToMid(n int) {
	if n == 1 {
		fmt.Println("move 1 from right to mid")
	}
	rightToLeft(n - 1)
	fmt.Println("move", n, "from right to mid")
	leftToMid(n - 1)
}

func rightToLeft(n int) {
	if n == 1 {
		fmt.Println("move 1 from right to left")
	}
	rightToMid(n - 1)
	fmt.Println("move", n, "from right to left")
	midToLeft(n - 1)
}

func midToLeft(n int) {
	if n == 1 {
		fmt.Println("movw 1 from mid to left")
	}
	midToRight(n - 1)
	fmt.Println("move", n, "from mid to left")
	rightToLeft(n - 1)
}

func midToRight(n int) {
	if n == 1 {
		fmt.Println("move 1 from mid to right")
	}
	midToLeft(n - 1)
	fmt.Println("move", n, "from mid to right")
	leftToRight(n - 1)
}

func Hano2(n int) {}

func move(from, to, other string, n int) {
	if n == 1 {
		fmt.Println("move 1 from", from, "to", to)
	} else {
		move(from, other, to, n-1)
		fmt.Println("move", n, "from", from, "to", to)
		move(other, to, from, n-1)
	}
}
