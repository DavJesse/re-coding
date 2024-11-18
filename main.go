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
	// Display greeting message and home menu
	greeting.Greeting()
	//greeting.HomeMenu()

	for {
		greeting.HomeMenu()
		scanner.Scan()

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

			// Append new task to tasks slice and print success message
			tasks = add.Append(task, tasks)
			fmt.Println("Task added successfully!")
			fmt.Println()
			continue // return to home

		} else if scanner.Text() == "2" {
			// Print all tasks
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
