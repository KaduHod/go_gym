package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type BuildExerciseService struct {
    db *sql.DB
}
func NewBuildExerciseService(db *sql.DB) *BuildExerciseService {
    return &BuildExerciseService{
        db: db,
    }
}
func (b *BuildExerciseService) Build(name string) error {
    // pegar descricao do exericicio
    exercise, err := b.getExercise(name)
    if err != nil {
        return err
    }
    // pegar relação de articulações, movimentos e musculos
    relations, err := b.getRelations()
    if err != nil {
        return err
    }
    jointsJson, err := json.Marshal(relations.Joints)
    if err != nil {
        fmt.Println(err)
        return err
    }
    musclesJson, err := json.Marshal(relations.MusclePortions)
    if err != nil {
        fmt.Println(err)
        return err
    }
    movementsJson, err := json.Marshal(relations.Movements)
    if err != nil {
        fmt.Println(err)
        return err
    }
    // Criar contexto da IA para processar exercicio e criar relacionamento com o banco
    prompt, err := b.prompt(exercise.Name, exercise.Description, string(jointsJson), string(musclesJson), string(movementsJson))
    fmt.Println(prompt)
    b.requestToAi(prompt)
    return nil
}
func (b *BuildExerciseService) getExercise(name string) (Exercise, error) {
    var exercise Exercise
    if err := b.db.QueryRow("SELECT id, name, description, info_link FROM exercise WHERE name = ?", name).Scan(&exercise.Id, &exercise.Name, &exercise.Description, &exercise.Link); err != nil {
        fmt.Println(err)
        return exercise, err
    }
    return exercise, nil
}
type MuscleGroup struct {
    Name string `json:"name"`
    Id int `json:"id"`
}
type MusclePortion struct {
    Name string `json:"name"`
    Id int `json:"id"`
}
type Movement struct {
    Name string `json:"name"`
    Id int `json:"id"`
}
type Joint struct {
    Name string `json:"name"`
    Id int `json:"id"`
}
type RelationsForBuild struct {
    Movements []Movement `json:"movements"`
    MuscleGroups []MuscleGroup `json:"muscle_groups"`
    MusclePortions []MusclePortion `json:"muscle_portions"`
    Joints []Joint `json:"joints"`
}
func (b *BuildExerciseService) getRelations() (RelationsForBuild, error) {
    var relations RelationsForBuild
    jointsRows, err := b.db.Query(`SELECT name, id FROM articulations`)
    if err != nil {
        fmt.Println(err)
        return relations, err
    }
    defer jointsRows.Close()
    var joints []Joint
    for jointsRows.Next() {
        var joint Joint
        if err := jointsRows.Scan(&joint.Name, &joint.Id); err != nil {
            fmt.Println(err)
            return relations, err
        }
        joints = append(joints, joint)
    }

    movementsRows, err := b.db.Query(`SELECT name, id FROM movements`)
    if err != nil {
        fmt.Println(err)
        return relations, err
    }
    defer movementsRows.Close()
    var movements []Movement
    for movementsRows.Next() {
        var movement Movement
        if err := movementsRows.Scan(&movement.Name, &movement.Id); err != nil {
            fmt.Println(err)
            return relations, err
        }
        movements = append(movements, movement)
    }
    muscleGroupsRows, err := b.db.Query(`SELECT name, id FROM muscle_group`)
    if err != nil {
        fmt.Println(err)
        return relations, err
    }
    defer muscleGroupsRows.Close()
    var muscleGroups []MuscleGroup
    for muscleGroupsRows.Next() {
        var muscleGroup MuscleGroup
        if err := muscleGroupsRows.Scan(&muscleGroup.Name, &muscleGroup.Id); err != nil {
            fmt.Println(err)
            return relations, err
        }
        muscleGroups = append(muscleGroups, muscleGroup)
    }
    musclePortionsRows, err := b.db.Query(`SELECT name, id FROM muscle_portion`)
    if err != nil {
        fmt.Println(err)
        return relations, err
    }
    defer musclePortionsRows.Close()
    var musclePortions []MusclePortion
    for musclePortionsRows.Next() {
        var musclePortion MusclePortion
        if err := musclePortionsRows.Scan(&musclePortion.Name, &musclePortion.Id); err != nil {
            fmt.Println(err)
            return relations, err
        }
        musclePortions = append(musclePortions, musclePortion)
    }
    relations.Joints = joints
    relations.Movements = movements
    relations.MuscleGroups = muscleGroups
    relations.MusclePortions = musclePortions
    return relations, nil
}
func (b *BuildExerciseService) prompt(exerciseName string, exerciseDescription string, jointsList string, musclesList string, movementsList string) (string, error) {
    skeleton, err := b.getPromptSkeleton()
    if err != nil {
        fmt.Println(err)
        return ``, err
    }
    skeleton = strings.ReplaceAll(skeleton, "{exercise_name}", exerciseName)
    skeleton = strings.ReplaceAll(skeleton, "{exercise_description}", exerciseDescription)
    skeleton = strings.ReplaceAll(skeleton, "{joints_list}", jointsList)
    skeleton = strings.ReplaceAll(skeleton, "{muscles_list}", musclesList)
    prompt := strings.ReplaceAll(skeleton, "{movements_list}", movementsList)
    return prompt, nil
}
func (b *BuildExerciseService) getPromptSkeleton() (string, error) {
    filePath := "/home/carlos/projetos/gym_golang/prompts/relacaoMusculosArticulacoesMovimentos.md"
    content, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println(err, "Erro ao pegar prompt skeleton")
        return "", err
    }
    return string(content), err
}
func (b *BuildExerciseService) requestToAi(content string) (string, error) {
    body := map[string]interface{}{
        "model": "deepseek-chat",
        "messages": []map[string]interface{}{
            {
                "role":    "user",
                "content": content,
            },
        },
    }
    bodyJson, err := json.Marshal(body)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    bodyBytes := bytes.NewBuffer(bodyJson)
    request, err := http.NewRequest("POST", "https://api.deepseek.com/v1/chat/completions", bodyBytes)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("Authorization", "Bearer <meu tokenzinho>")
    requestClient := &http.Client{}
    response, err := requestClient.Do(request)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    defer response.Body.Close()
    responseBody, err := io.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    fmt.Println(string(responseBody))
    fmt.Println(response.StatusCode)
    return "", nil
}
