package remove

import "davodhiambo/models"

func DeleteTask(id int, tasks []models.Task) []models.Task {
	var result []models.Task

	for _, task := range tasks {
		if task.ID != id {
			result = append(result, task)
		}
	}
	return result
}
