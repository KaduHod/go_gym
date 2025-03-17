package services

import (
	"database/sql"
	"fmt"
)

type ListExercisesService struct {
    db *sql.DB
}
type MMJ struct {
    RelationId int `json:"relation_id"`
    JointName string `json:"join_name"`
    MovimentName string `json:"moviment_name"`
    MusclePortionName string `json:"muscle_portion_name"`
    MuscleGroupName string `json:"muscle_group_name"`
    Role string `json:"role"`
}
func NewListExercisesService(db *sql.DB) *ListExercisesService {
    return &ListExercisesService{
        db: db,
    }
}
func (s *ListExercisesService) GetById(id int) (Exercise, error) {
    var exercise Exercise
    err := s.db.QueryRow("SELECT id, name, description, info_link FROM exercise WHERE id = ?", id).Scan(&exercise.Id, &exercise.Name, &exercise.Description, &exercise.Link)
    return exercise, err
}
func (s *ListExercisesService) ListExercises() ([]Exercise, error) {
    rows, err := s.db.Query("SELECT id, name FROM exercise")
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    defer rows.Close()
    var exercises []Exercise
    for rows.Next() {
        var exercise Exercise
        if err := rows.Scan(&exercise.Id,&exercise.Name); err != nil {
            return nil, err
        }
        exercises = append(exercises, exercise)
    }
    return exercises, nil
}
func (s *ListExercisesService) ListExercisesJmmByRole(exercise Exercise) (map[string][]MMJ, error) {
    rows, err := s.db.Query(`
SELECT
	amm.id as relation_id,
	a.name as join_name,
	m.name as moviment_name,
	mp.name as muscle_portion_name,
	mg.name as muscle_group_name,
    eamm.role as role
FROM exercise_amm eamm
INNER JOIN exercise e on e.id = eamm.exercise_id
INNER JOIN articulation_movement_muscle amm on amm.id = eamm.amm_id
INNER JOIN articulations a on a.id = amm.articulation_id
INNER JOIN muscle_portion mp on mp.id = amm.muscle_portion_id
INNER JOIN muscle_group mg on mp.muscle_group_id = mg.id
INNER JOIN movements m on m.id = amm.movement_id
WHERE eamm.exercise_id  = ?`, exercise.Id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var mmjs []MMJ
    for rows.Next() {
        var mmj MMJ
        if err := rows.Scan(&mmj.RelationId, &mmj.JointName, &mmj.MovimentName, &mmj.MusclePortionName, &mmj.MuscleGroupName, &mmj.Role); err != nil {
            return nil, err
        }
        mmjs = append(mmjs, mmj)
    }
    var mapByRole = make(map[string][]MMJ)
    mapByRole["stabilizer"] = make([]MMJ, 0)
    mapByRole["eccentric"] = make([]MMJ, 0)
    mapByRole["concentric"] = make([]MMJ, 0)
    for _, mmj := range mmjs {
        mapByRole[mmj.Role] = append(mapByRole[mmj.Role], mmj)
    }
    return mapByRole, nil
}
func (s *ListExercisesService) ListExerciseJmm(exercise Exercise) ([]MMJ, error) {
    rows, err := s.db.Query(`
SELECT
	amm.id as relation_id,
	a.name as join_name,
	m.name as moviment_name,
	mp.name as muscle_portion_name,
	mg.name as muscle_group_name
FROM exercise_amm eamm
INNER JOIN exercise e on e.id = eamm.exercise_id
INNER JOIN articulation_movement_muscle amm on amm.id = eamm.amm_id
INNER JOIN articulations a on a.id = amm.articulation_id
INNER JOIN muscle_portion mp on mp.id = amm.muscle_portion_id
INNER JOIN muscle_group mg on mp.muscle_group_id = mg.id
INNER JOIN movements m on m.id = amm.movement_id
WHERE eamm.exercise_id  = ?`, exercise.Id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var mmjs []MMJ
    for rows.Next() {
        var mmj MMJ
        if err := rows.Scan(&mmj.RelationId, &mmj.JointName, &mmj.MovimentName, &mmj.MusclePortionName, &mmj.MuscleGroupName); err != nil {
            return nil, err
        }
        mmjs = append(mmjs, mmj)
    }
    return mmjs, nil
}
