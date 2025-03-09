package tests

import (
	"kaduhod/gym/app/services"
	"kaduhod/gym/database"
	"testing"
)

func TestCreateExercise(t *testing.T) {
    service := services.NewCreateExerciseService(database.ConnetionMysql())
    _, err := service.Create("Teste", "Teste")
    if err != nil {
        t.Errorf("Error creating exercise: %v", err)
        return
    }
}
func TestCreateExerciseCapital(t *testing.T) {
    service := services.NewCreateExerciseService(database.ConnetionMysql())
    _, err := service.Create("TESTE2", "TESTE")
    if err != nil {
        t.Errorf("Error creating exercise: %v", err)
        return
    }
}
