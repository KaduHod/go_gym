
package tests

import (
	"kaduhod/gym/app/services"
	"kaduhod/gym/database"
	"os"
	"testing"
)

var service *services.ListExercisesService

// TestMain configura o ambiente antes de rodar os testes
func TestMain(m *testing.M) {
    // Configuração antes dos testes
    db := database.ConnetionMysql()
    service = services.NewListExercisesService(db)

    // Executa os testes
    exitCode := m.Run()

    // Finaliza testes (caso precise limpar algo)
    os.Exit(exitCode)
}
func TestGetExerciseById(t *testing.T) {
    result, err := service.GetById(87)
    if err != nil {
        t.Errorf("Error listing exercises: %v", err)
        return
    }
    t.Logf("Exercises: %v", result)
}
func TestListExercises(t *testing.T) {
    result, err := service.ListExercises()
    if err != nil {
        t.Errorf("Error listing exercises: %v", err)
        return
    }
    t.Logf("Exercises: %v", result)
}

func TestMMJFromExercise(t *testing.T) {
    result, err := service.ListExerciseJmm(services.Exercise{Id: 87})
    if err != nil {
        t.Errorf("Error listing exercises: %v", err)
        return
    }
    t.Logf("Exercises: %v", result)
}

