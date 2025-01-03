package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"davodhiambo/add"
	"davodhiambo/display"
	"davodhiambo/greeting"
	"davodhiambo/models"
	"davodhiambo/remove"
	"davodhiambo/sorting"
)

func main() {
	var tasks []models.Task
	var task models.Task
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	// Display greeting message and home menu
	greeting.Greeting()
	// greeting.HomeMenu()

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
			if len(tasks) == 0 {
				fmt.Println("You have no tasks")
				fmt.Println()
				continue
			}
			sorting.Sort(tasks)
			fmt.Println("Your Tasks:")
			display.DisplayTasks(tasks)

		} else if scanner.Text() == "3" {
			if len(tasks) == 0 {
				fmt.Println("You have no tasks")
				fmt.Println()
				continue
			}
			fmt.Println("Which of these tasks do you wish to delete?")
			display.DisplayTasks(tasks)
			fmt.Println()
			fmt.Println("Key in the ID of the task you wish to delete:")
			scanner.Scan()

			id, err := strconv.Atoi(scanner.Text())

			for err != nil || id > len(tasks) || id < 1 {
				if err != nil {
					fmt.Println("Invalid input.")
				} else {
					fmt.Printf("Task %d does not exist.\n", id)
				}
				fmt.Println("Try again")
				fmt.Println()
				fmt.Println("Which of these tasks do you wish to delete?")
				display.DisplayTasks(tasks)
				fmt.Println()
				fmt.Println("Key in the ID of the task you wish to delete:")
				scanner.Scan()
				id, err = strconv.Atoi(scanner.Text())
			}

			tasks = remove.DeleteTask(id, tasks)
			fmt.Printf("Task %d deleted successfully!\n\n", id)
			continue

		} else if scanner.Text() == "4" {
			if len(tasks) == 0 {
				fmt.Println("You have no tasks")
				fmt.Println()
				continue
			}
			fmt.Println("Which of these tasks do you wish to update?")
			display.DisplayTasks(tasks)
			fmt.Println()
			fmt.Println("Key in the ID of the task you wish to update:")
			scanner.Scan()

			id, err := strconv.Atoi(scanner.Text())

			for err != nil || id > len(tasks) || id < 1 {
				if err != nil {
					fmt.Println("Invalid input.")
				} else {
					fmt.Printf("Task %d does not exist.\n", id)
				}
				fmt.Println("Try again")
				fmt.Println()
				fmt.Println("Which of these tasks do you wish to update?")
				display.DisplayTasks(tasks)
				fmt.Println()
				fmt.Println("Key in the ID of the task you wish to update:")
				scanner.Scan()
				id, err = strconv.Atoi(scanner.Text())
			}


		} else if scanner.Text() == "5" {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("Wrong Input\nTry Again")
			continue
		}
	}
}
