package tests

import (
	"kaduhod/gym/app/services"
	"kaduhod/gym/database"
	"testing"
)

func TestGetExercises(t *testing.T) {
    service := services.NewGetExerciseService(database.ConnetionMysql())
    _, err := service.GetByName("Bench Press")
    if err != nil {
        t.Errorf("Error listing exercise: %v", err)
        return
    }
}
