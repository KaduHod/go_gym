package services

import (
	"fmt"
	"slices"
)

type ExerciseStatsService struct {
    exercise Exercise
}
type ExerciseStats struct {
    JointsCount int
    MovementsCount int
    MusclesCount int
    MusclePortionsCount int
    Joints []string
    Movements []string
    MusclesGroups map[string][]string
}
func NewExerciseStatsService(exercise Exercise) *ExerciseStatsService {
    return &ExerciseStatsService{
        exercise: exercise,
    }
}

func (self *ExerciseStatsService) GetStats(relations []MMJ) ExerciseStats {
    var joints []string
    var movements []string
    musclesGroups := make(map[string][]string)
    var musclePortions int
    for _, item := range relations {
        if !slices.Contains(joints, item.JointName) {
            joints = append(joints, item.JointName)
        }
        if !slices.Contains(movements, item.MovimentName) {
            movements = append(movements, item.MovimentName)
        }
        if _, exists := musclesGroups[item.MuscleGroupName]; !exists {
            musclesGroups[item.MuscleGroupName] = []string{}
        }
        if !slices.Contains(musclesGroups[item.MuscleGroupName], item.MusclePortionName) {
            musclesGroups[item.MuscleGroupName] = append(musclesGroups[item.MuscleGroupName], item.MusclePortionName)
            musclePortions++
        }
    }
    stats := ExerciseStats {
        JointsCount: len(joints),
        MovementsCount: len(movements),
        MusclesCount: len(musclesGroups),
        MusclePortionsCount: musclePortions,
        Joints: joints,
        Movements: movements,
        MusclesGroups: musclesGroups,
    }
    fmt.Println(stats)
    return stats
}
