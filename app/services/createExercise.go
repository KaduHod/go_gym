package services

import "database/sql"

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
func (self *CreateExerciseService) Create() ([]string, error) {
    return nil, nil
}
