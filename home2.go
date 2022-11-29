package main

import (
	"fmt"
)

func transformation(a []int) {

	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Исходнк", a)
	fmt.Println("Переворот")
	transformation(a)
	fmt.Println(a)

}
