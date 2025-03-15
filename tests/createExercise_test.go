package tests

import (
	"kaduhod/gym/app/services"
	"testing"
)

func TestCreateExercise(t *testing.T) {
    if db == nil {
        t.Errorf("Database not initialized")
        return
    }
    service := services.NewCreateExerciseService(db)
    _, err := service.Create("Teste", "Teste", "teste")
    if err != nil {
        t.Errorf("Error creating exercise: %v", err)
        return
    }
}
func TestCreateExerciseCapital(t *testing.T) {
    service := services.NewCreateExerciseService(db)
    _, err := service.Create("TESTE2", "TESTE", "teste")
    if err != nil {
        t.Errorf("Error creating exercise: %v", err)
        return
    }
}
func TestUpdateExercise(t *testing.T) {
    service := services.NewCreateExerciseService(db)
    _, err := service.Update(services.Exercise{Id: 178, Name: "Bench Press Teste", Description: "The bench press is a weight training exercise where a person presses a weight upward while lying on a bench, primarily targeting the pectoralis major, anterior deltoids, and triceps brachii. Stabilizing muscles in the back, legs, and core also assist. Typically performed with a barbell or dumbbells, it is a key lift in powerlifting alongside the squat and deadlift and the sole lift in Paralympic powerlifting. Widely used in bodybuilding and strength training, the bench press enhances upper body strength, power, and muscle development.", Link: "teste"})
    if err != nil {
        t.Errorf("Error updating exercise: %v", err)
        return
    }
}
