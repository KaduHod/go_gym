package services

import (
	"database/sql"
	"errors"
	"kaduhod/gym/app/utils"
)

type Exercise struct {
    Name string `json:"name" validate:"required"`
    Description string `json:"description" validate:"required"`
    Id int
    Link string
}
type CreateExerciseService struct {
    db *sql.DB
}
func NewCreateExerciseService(db *sql.DB) *CreateExerciseService {
    return &CreateExerciseService{
        db: db,
    }
}
func (self *CreateExerciseService) Create(name string, description string) ([]string, error) {
    exists, err := self.exerciseExists(name)
    if exists {
        return []string{"Exercise already exists"}, errors.New("Exercise already exists")
    }
    if err != nil {
        return nil, err
    }
    _, err = self.db.Exec("INSERT INTO exercise (name, description) VALUES (?, ?)", utils.CapitalizePhrase(name), description)
    if err != nil {
        return nil, err
    }
    return nil, nil
}
func (self *CreateExerciseService) exerciseExists(name string) (bool, error) {
    var exists bool
    err := self.db.QueryRow("SELECT EXISTS(SELECT * FROM exercise WHERE name = ?)", utils.CapitalizePhrase(name)).Scan(&exists)
    return exists, err
}
