package testing

import (
	"testing"

	"davodhiambo/add"
	"davodhiambo/models"
)

func TestAppend(t *testing.T) {
	subject := models.Task{Description: "Go to Work", Priority: 3}
	got := add.Append(subject, []models.Task{})
	subject.ID = 1
	expect := []models.Task{subject}

	if got[0].ID != expect[0].ID || got[0].Description != expect[0].Description || got[0].Priority != expect[0].Priority {
		t.Errorf("Expected: %v, Got: %v", expect[0], got[0])
		t.Errorf("Test failed!")
	}
}
