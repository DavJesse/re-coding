package greeting

import "fmt"

func Greeting() {
	fmt.Println("Welcome to Your To-Do List\nPlease Pick an Option")
	fmt.Println()
}

func HomeMenu() {
	fmt.Println("Key '1' to Add Task")
	fmt.Println("Key '2' to View Task")
	fmt.Println("Key '3' to Delete Task")
	fmt.Println("Key '4' to Update Task")
	fmt.Println("Key '5' to Exit Program")
	fmt.Println()
}
