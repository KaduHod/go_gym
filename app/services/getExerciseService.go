package services

import (
	"database/sql"
	"kaduhod/gym/app/utils"
)

type GetExerciseService struct {
    db *sql.DB
}

func NewGetExerciseService(db *sql.DB) *GetExerciseService {
    return &GetExerciseService{
        db: db,
    }
}

func (self *GetExerciseService) GetByName(name string) (Exercise, error) {
    var exercise Exercise
    err := self.db.QueryRow("SELECT id, name, description FROM exercise WHERE name = ?", utils.CapitalizePhrase(name)).Scan(&exercise.Id, &exercise.Name, &exercise.Description)
    return exercise, err
}
