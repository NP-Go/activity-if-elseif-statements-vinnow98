package main

import "fmt"

func activity2() {
	number1 := 17
	resultMessage := "No outcome."
	if number1%2 == 0 {
		fmt.Print("This is an even number")
	} else if number1%2 == 1 {
		fmt.Println("This is an odd number")
	} else {
		fmt.Println(resultMessage)

	}
	if number1 > 9 {
		fmt.Println("This number has more than 1 digit")
	} else if number1 < -9 {
		fmt.Println("This number has more than 1 digit")
	} else {
		fmt.Println("This number has only 1 digit")
	}
}

func main() {
	activity2()
}
