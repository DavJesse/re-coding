package main

import (
	"bufio"
	"fmt"
	"os"

	"davodhiambo/add"
	"davodhiambo/greeting"
	"davodhiambo/models"
)

func main() {
	var tasks []models.Task
	var task models.Task
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	greeting.Greeting()
	for scanner.Scan() {
		fmt.Println("Key '1' to Add Task")
		fmt.Println("Key '2' to View Task")
		fmt.Println("Key '3' to Delete Task")
		fmt.Println("Key '4' to Update Task")
		fmt.Println("Key '5' to Exit Program")
		fmt.Println()
		
		if scanner.Text() == "1" {
			// Prompt user to add a description for new task
			add.PromptDescription()
			scanner.Scan()
			task = add.AddDescription(scanner.Text(), task)

			// Prompt user to add a priority for new task
			add.AddPriority()
			scanner.Scan()
			task, err = add.CheckPriority(scanner.Text(), task)
		
			// Handle errors when adding priority
			for err != nil {
				fmt.Println(err)
				add.AddPriority()
				scanner.Scan()
				task, err = add.CheckPriority(scanner.Text(), task)
			}
			tasks = add.Append(task, tasks)
			fmt.Println("Task added successfully!")
			fmt.Println()
			continue
		} else if scanner.Text() == "2" {
			fmt.Println("Your Tasks:")
			for _, task := range tasks {
				fmt.Println(task.ID)
				fmt.Println(task.Description)
				fmt.Println(task.Priority)
			}
		} else if scanner.Text() == "3" {
		} else if scanner.Text() == "4" {
		} else if scanner.Text() == "5" {
		} else {
			fmt.Println("Wrong Input\nTry Again")
			continue
		}
	}
}
