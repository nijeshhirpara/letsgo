package models

import "gorm.io/gorm"

// Task table contains information about each user
type Task struct {
	gorm.Model
	ID     uint
	Name   string
	Status string
	TeamID uint
	Team   Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

// TaskRepository interface
type TaskRepository interface {
	CreateTask(t Team, u User, task Task) error
	ListTasksByTeam(teamID uint) (tasks []Task)
}
