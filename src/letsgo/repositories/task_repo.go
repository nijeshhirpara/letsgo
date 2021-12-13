package repositories

import (
	"letsgo/models"
	"log"

	"gorm.io/gorm"
)

// TaskRepo implements models.TaskRepository interface
type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

// CreateTask creats a task
func (taskRepo *TaskRepo) CreateTask(team models.Team, user models.User, task models.Task) error {
	task.Team = team
	task.User = user
	res := taskRepo.db.Create(&task)
	return res.Error
}

func (taskRepo *TaskRepo) ListTasksByTeam(teamID uint) (tasks []models.Task) {
	result := taskRepo.db.Preload("Team").Preload("User").Find(&tasks)

	if result.Error != nil {
		log.Println(result.Error.Error())
	}

	return
}
