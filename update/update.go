package update

import "davodhiambo/models"

func UpdateTask(id int, description string, priority int, tasks []models.Task) []models.Task {
	for _, task := range tasks {
		if task.ID == id {
			if description != "" {
				task.Description = description
			}
			if priority > 0 && priority <= 3 {
				task.Priority = priority
			}
		}
	}
	return tasks
}
