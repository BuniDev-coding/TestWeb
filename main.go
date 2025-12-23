package main

import "fmt"

func main() {
	var fname [3]string = [3]string{"John", "Jane", "Doe"}
	fmt.Println("First Name:", fname[0])
	fmt.Println("Second Name:", fname[1])
	fmt.Println("Third Name:", fname[2])


	var input int
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter an integer: ")
	fmt.Scanln(&input)
	if input >= 50 && input <= 55 {
		fmt.Println(name, "ได้ D ")
	} else if input >= 56 && input <= 59 {
		fmt.Println(name, "ได้ D+ ")
	} else if input >= 60 && input <= 64 {
		fmt.Println(name, "ได้ C ")
	} else if input >= 65 && input <= 69 {
		fmt.Println(name, "ได้ C+ ")
	} else if input >= 70 && input <= 74 {
		fmt.Println(name, "ได้ B ")
	} else if input >= 75 && input <= 79 {
		fmt.Println(name, "ได้ B+ ")
	} else if input >= 80 {
		fmt.Println(name, "ได้ A ")
	} else {
		fmt.Println("Invalid score")
		return
	}
}
