package sort

import (
	"davodhiambo/models"
	"sort"
)

func Sort(tasks []models.Task) []models.Task {
	sort.Slice(tasks, func (i, j int) bool {
		return tasks[i].Priority < tasks[j].Priority
	})
	return tasks
}

