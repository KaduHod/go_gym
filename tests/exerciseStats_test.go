package tests

import (
	"fmt"
	"testing"
)

func TestGetExerciseStats(t *testing.T) {
    benchPress, err := service.GetByName("Bench Press")
    if err != nil {
        t.Errorf("Error listing exercise by name: %v", err)
        return
    }
    mmj, err := service.ListExerciseJmm(benchPress)
    if err != nil {
        t.Errorf("Error listing exercise mmj: %v", err)
        return
    }
    stats := exerciseStatsService.GetStats(mmj)
    fmt.Println(stats)
}
