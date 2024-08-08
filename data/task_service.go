package data

import (
	"time"

	"github.com/amha-mersha/go_taskmanager/models"
)

type TaskError struct {
	message string
}

func (err TaskError) Error() string {
	return err.message
}

const (
	IDNotFound        = "No item found with the specified ID."
	TaskAlreadyExists = "The task already exists in the database"
	MalformedJSON     = "Sent a malfomed JSON."
	MismatchedFormat  = "The task have a mismatched stucture."
	MissingRequireds  = "There are some missing required feilds."
)

var tasks = map[int64]models.Task{
	0: {
		ID:          0,
		Title:       "Prepare presentation",
		Description: "Prepare slides for the upcoming presentation",
		Status:      "pending",
		Priority:    "",
		DueDate:     time.Time{},
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	},
	1: {
		ID:          1,
		Title:       "Write unit tests",
		Description: "Write unit tests for the new features implemented",
		Status:      "pending",
		Priority:    "",
		DueDate:     time.Time{},
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	},
	2: {
		ID:          2,
		Title:       "Complete project report",
		Description: "Finish the annual project report by the end of the week",
		Status:      "completed",
		Priority:    "high",
		DueDate:     time.Time{},
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	},
	3: {
		ID:          3,
		Title:       "Team meeting",
		Description: "Attend the weekly team meeting on Friday",
		Status:      "completed",
		Priority:    "",
		DueDate:     time.Time{},
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	},
}

func GetTasks() map[int64]models.Task {
	return tasks
}

func GetTaskByID(taskID int64) (models.Task, error) {
	if _, exist := tasks[taskID]; !exist {
		return models.Task{}, TaskError{message: IDNotFound}
	}
	return tasks[taskID], nil
}

func UpdateTask(taskID int64, updatedTask models.Task) error {
	if _, exist := tasks[taskID]; !exist {
		return TaskError{message: IDNotFound}
	}
	tasks[taskID] = updatedTask
	return nil
}

func DeleteTask(taskID int64) error {
	if _, exist := tasks[taskID]; !exist {
		return TaskError{message: IDNotFound}
	}
	delete(tasks, taskID)
	return nil
}

func PostTask(newTask models.Task) error {
	if _, exist := tasks[int64(newTask.ID)]; exist {
		return TaskError{message: TaskAlreadyExists}
	}
	tasks[int64(newTask.ID)] = newTask
	return nil
}
