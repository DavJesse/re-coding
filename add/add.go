package add

import (
	"davodhiambo/models"
	"errors"
	"fmt"
)

func Append(task models.Task, tasks []models.Task) []models.Task {
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
	return tasks
}

func PromptDescription() {
	fmt.Println("\nDescribe your task")
}

func AddPriority() {
	fmt.Println("\nHow would you define Priority for this task?")
	fmt.Println("Key '1' for High Priority")
	fmt.Println("Key '2' for Medium Priority")
	fmt.Println("Key '3' for Low Priority")
	//fmt.Println("Key '4' to go back to Home")
}

func CheckPriority(scan string, task models.Task) (models.Task, error) {
	err := errors.New("Wrong Input\nTry Again")
	
	if scan == "1" {
		task.Priority = 1
		return task, nil
	} else if scan == "2" {
		task.Priority = 2
		return task, nil
	} else if scan == "3" {
		task.Priority = 3
		return task, nil
	} else {
		return models.Task{}, err
	}
}

func AddDescription(scan string, task models.Task) models.Task{
	task.Description = scan
	return task
}