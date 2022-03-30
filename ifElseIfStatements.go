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
func activity3() {
	var input int
	fmt.Println("Input a number!")
	fmt.Scanln(&input)
	a := input%5 == 0
	b := input%6 == 0
	if a && b {
		fmt.Println("This number is divisible by 5 and 6")
	} else {
		fmt.Println("This number is not divisible by 5 and 6")
	}

}
func main() {
	activity2()
	activity3()
}
