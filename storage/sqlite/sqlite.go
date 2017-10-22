package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jvikstedt/watchful/model"
	"github.com/jvikstedt/watchful/storage"
)

func New(filepath string) (storage.Service, error) {
	db, err := gorm.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	storage := &sqlite{
		db: db,
	}

	storage.EnsureTables()

	return storage, nil
}

type sqlite struct {
	db *gorm.DB
}

func (s *sqlite) JobCreate(job *model.Job) error {
	return s.db.Create(job).Error
}

func (s *sqlite) TaskCreate(task *model.Task) error {
	return s.db.Create(task).Error
}

func (s *sqlite) TaskDelete(id int) error {
	task := model.Task{ID: uint(id)}
	return s.db.Delete(&task).Error
}

func (s *sqlite) TasksByJobID(jobID int, tasks *[]model.Task) error {
	job := model.Job{ID: uint(jobID)}
	return s.db.Model(&job).Related(tasks).Error
}

func (s *sqlite) EnsureTables() error {
	return s.db.AutoMigrate(
		&model.Job{},
		&model.Task{},
		&model.Input{}).Error
}

func (s *sqlite) Close() error {
	return s.db.Close()
}
