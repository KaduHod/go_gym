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
    Link string `json:"link" validate:"required"`
}
type CreateExerciseService struct {
    db *sql.DB
}
func NewCreateExerciseService(db *sql.DB) *CreateExerciseService {
    return &CreateExerciseService{
        db: db,
    }
}
func (self *CreateExerciseService) Update(exercise Exercise) ([]string, error) {
    _, err := self.db.Exec("UPDATE exercise SET name = ?, description = ?, info_link = ? WHERE id = ?", utils.CapitalizePhrase(exercise.Name), exercise.Description, exercise.Link, exercise.Id)
    if err != nil {
        return nil, err
    }
    return nil, nil
}
func (self *CreateExerciseService) Create(name string, description string, link string) ([]string, error) {
    exists, err := self.exerciseExists(name)
    if exists {
        return []string{"Exercise already exists"}, errors.New("Exercise already exists")
    }
    if err != nil {
        return nil, err
    }
    _, err = self.db.Exec("INSERT INTO exercise (name, description, info_link) VALUES (?, ?, ?)", utils.CapitalizePhrase(name), description, link)
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
