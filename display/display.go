package display

import (
	"fmt"

	"davodhiambo/models"
)

func DisplayTasks(tasks []models.Task) {
	for _, task := range tasks {
		fmt.Printf("Task: %d\n\n", task.ID)
		fmt.Printf("Description:\n %s\n\n", task.Description)
		fmt.Printf("Priority: %d\n\n", task.Priority)
	}
}
