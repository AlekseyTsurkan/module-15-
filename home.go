package main

import "fmt"

func main() {
	sum := 0
	nesum := 0
	var number [10]int

	for i := 0; i < 10; i++ {
		fmt.Println("Введите число ", i+1)
		fmt.Scan(&number[i])

	}
	for _, number := range number {

		if number%2 != 0 {
			// fmt.Println("не четные числа",i)

			nesum += 1
		} else {
			//fmt.Println("Число четное", nesum)
			sum += 1
		}
	}
	fmt.Println("Числа четные", sum, "И числа не четных ", nesum)
}
