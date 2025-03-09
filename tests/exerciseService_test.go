
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
    after()
    os.Exit(exitCode)
}
func after() {
    db := database.ConnetionMysql()
    db.Exec("DELETE FROM user WHERE email = ?", "test@mail2.com")
    db.Exec("DELETE FROM exercise WHERE name = 'Teste'")
    db.Exec("DELETE FROM exercise WHERE name = 'Teste2'")
}
func TestGetExerciseById(t *testing.T) {
    _, err := service.GetById(87)
    if err != nil {
        t.Errorf("Error listing exercises: %v", err)
        return
    }
}
func TestListExercises(t *testing.T) {
    _, err := service.ListExercises()
    if err != nil {
        t.Errorf("Error listing exercises: %v", err)
        return
    }
}

func TestMMJFromExercise(t *testing.T) {
    _, err := service.ListExerciseJmm(services.Exercise{Id: 87})
    if err != nil {
        t.Errorf("Error listing exercises: %v", err)
        return
    }
}

