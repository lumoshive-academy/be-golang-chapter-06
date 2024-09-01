package main

import "fmt"

func main() {
	var a, b int = 10, 20
	fmt.Println("Penjumlahan:", a+b)
	fmt.Println("Pengurangan:", a-b)
	fmt.Println("Perkalian:", a*b)
	fmt.Println("Pembagian:", a/b)
	fmt.Println("Modulus:", a%b)

	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)

	var c, d bool = true, false
	fmt.Println("a && b:", c && d)
	fmt.Println("a || b:", c || d)
	fmt.Println("!a:", !c)

}
