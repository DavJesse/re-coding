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
		if scanner.Text() == "1" {
			add.PromptDescription()
			scanner.Scan()
			task = add.AddDescription(scanner.Text(), task)

			add.AddPriority()
			scanner.Scan()
			task, err = add.CheckPriority(scanner.Text(), task)

			for err != nil {
				fmt.Println(err)
				add.AddPriority()
				scanner.Scan()
				task, err = add.CheckPriority(scanner.Text(), task)
			}
			tasks = add.Append(task, tasks)
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
