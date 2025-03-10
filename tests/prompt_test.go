package tests

import (
	"kaduhod/gym/app/services"
	"kaduhod/gym/database"
	"testing"
)

func TestPrompt(t *testing.T) {
    db := database.ConnetionMysql()
    service := services.NewBuildExerciseService(db)
    if err := service.Build("Bench Press Teste"); err != nil {
        t.Error(err)
    }
    // TODO: implement this
}
