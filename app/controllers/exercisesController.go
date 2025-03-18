package controllers

import (
	"fmt"
	"kaduhod/gym/app/services"
	"kaduhod/gym/app/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ExercisesController struct {
    Controller
    ListExercisesService services.ListExercisesService
    CreateExerciseService services.CreateExerciseService
    BuildExerciseService services.BuildExerciseService
}
func NewExerciseController(listExercisesService services.ListExercisesService, createExerciseService services.CreateExerciseService, buildExerciseService services.BuildExerciseService) *ExercisesController {
    return &ExercisesController{
        ListExercisesService: listExercisesService,
        CreateExerciseService: createExerciseService,
        BuildExerciseService: buildExerciseService,
    }
}
type ExerciseViewData struct {
    Exercises []services.Exercise
}
func (controller *ExercisesController) ListExercises(c echo.Context) error {
    exercises, err := controller.ListExercisesService.ListExercises()
    if err != nil {
        return c.String(400, err.Error())
    }
    var bechPess services.Exercise
    for _, exercise := range exercises {
        if exercise.Name == "Bench Press" {
            bechPess = exercise
        }
    }
    if bechPess.Id == 0 {
        if err := c.Render(200, "list_exercises2", map[string]interface{}{"Exercises": exercises}); err != nil {
            fmt.Println(err.Error())
            return err
        }
    }
    bechPess, err = controller.ListExercisesService.GetById(bechPess.Id)
    if err != nil {
        if err := c.Render(200, "list_exercises2", map[string]interface{}{"Exercises": exercises}); err != nil {
            fmt.Println(err.Error())
            return err
        }
    }
    mmjs, err := controller.ListExercisesService.ListExercisesJmmByRole(bechPess)
    if err != nil {
        if err := c.Render(200, "list_exercises2", map[string]interface{}{"Exercises": exercises}); err != nil {
            fmt.Println(err.Error())
            return err
        }
    }
    var mmjsForStats []services.MMJ
    for _, items := range mmjs {
        for _, item := range items {
            mmjsForStats = append(mmjsForStats, item)
        }
    }
    statsService := services.NewExerciseStatsService(bechPess)
    stats := statsService.GetStats(mmjsForStats)
    if err := c.Render(200, "list_exercises2", map[string]interface{}{
        "Exercises": exercises,
        "Mmjs": mmjs,
        "ImageURL": "public/images/exercises/"+strconv.Itoa(bechPess.Id)+".jpeg",
        "Exercise": bechPess,
        "Stats": stats,
    }); err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}
type MMJViewData struct {
    Mmjs []services.MMJ
    ImageURL string
    Exercise services.Exercise
}
func (controller *ExercisesController) ExerciseDetail(c echo.Context) error {
    exerciseId := c.Param("id")
    if exerciseId == "" {
        return c.String(400, "id is required")
    }
    exerciseIdInt, err := strconv.Atoi(exerciseId)
    if err != nil {
        return c.String(400, err.Error())
    }
    exercise, err := controller.ListExercisesService.GetById(exerciseIdInt)
    if err != nil {
        return c.String(400, err.Error())
    }
    mmjs, err := controller.ListExercisesService.ListExercisesJmmByRole(exercise)
    if err != nil {
        return c.String(400, err.Error())
    }
    var mmjsForStats []services.MMJ
    for _, items := range mmjs {
        for _, item := range items {
            mmjsForStats = append(mmjsForStats, item)
        }
    }
    statsService := services.NewExerciseStatsService(exercise)
    fmt.Println(statsService.GetStats(mmjsForStats))
    if err := c.Render(200, "exercise_page", map[string]interface{}{
       "Mmjs": mmjs,
       "ImageURL": "public/images/exercises/"+strconv.Itoa(exercise.Id)+".jpeg",
       "Exercise": exercise,
       "Stats": statsService.GetStats(mmjsForStats),
    }); err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}
func (controller *ExercisesController) EditExerciseForm(c echo.Context) error {
    exerciseId := c.Param("id")
    if exerciseId == "" {
        return c.String(400, "id is required")
    }
    exerciseIdInt, err := strconv.Atoi(exerciseId)
    if err != nil {
        return c.String(400, err.Error())
    }
    exercise, err := controller.ListExercisesService.GetById(exerciseIdInt)
    if err != nil {
        return c.String(400, err.Error())
    }
    mmjs, err := controller.ListExercisesService.ListExercisesJmmByRole(exercise)
    if err != nil {
        return c.String(400, err.Error())
    }
    if err := c.Render(200, "edit_exercise", map[string]interface{}{
        "Mmjs": mmjs,
        "ImageURL": "public/images/exercises/"+strconv.Itoa(exercise.Id)+".jpeg",
        "Exercise": exercise,
    }); err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}
func (controller *ExercisesController) EditExercise(c echo.Context) error {
    name := c.FormValue("name")
    description := c.FormValue("description")
    link := c.FormValue("link")
    id := c.FormValue("id")
    idInt, err := strconv.Atoi(id)
    fmt.Println(idInt, "AQUIII")
    exercise, err := controller.ListExercisesService.GetById(idInt)
    if err != nil {
        return c.String(400, err.Error())
    }
    mmjs, err := controller.ListExercisesService.ListExerciseJmm(exercise)
    if err != nil {
        return c.String(400, err.Error())
    }
    exercise.Name = name
    exercise.Description = description
    exercise.Link = link
    exercise.Id = idInt
    messages, err := controller.CreateExerciseService.Update(exercise)
    if err != nil {
        fmt.Println(err.Error())
        return c.Render(400, "edit_exercise", map[string]interface{}{
            "Errors": messages,
        })
    }
    viewData := MMJViewData{
        Mmjs: mmjs,
        ImageURL: "public/images/exercises/"+strconv.Itoa(exercise.Id)+".jpeg",
        Exercise: exercise,
    }
    if err := c.Render(200, "edit_exercise", viewData); err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}
func (controller *ExercisesController) Build(c echo.Context) error {
    exerciseId := c.Param("id")
    if exerciseId == "" {
        return c.String(400, "id is required")
    }
    exerciseIdInt, err := strconv.Atoi(exerciseId)
    if err != nil {
        return c.String(400, err.Error())
    }
    exercise, err := controller.ListExercisesService.GetById(exerciseIdInt)
    if err != nil {
        return c.String(400, err.Error())
    }
    if err := controller.BuildExerciseService.Build(exercise.Name); err != nil {
        return c.String(400, err.Error())
    }
    mmjs, err := controller.ListExercisesService.ListExerciseJmm(exercise)
    if err != nil {
        return c.String(400, err.Error())
    }
    viewData := MMJViewData{
        Mmjs: mmjs,
        ImageURL: "public/images/exercises/"+strconv.Itoa(exercise.Id)+".jpeg",
        Exercise: exercise,
    }
    if err := c.Render(200, "edit_exercise", viewData); err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}
func (controller *ExercisesController) CreateExerciseForm(c echo.Context) error {
    exercises, err := controller.ListExercisesService.ListExercises()
    if err != nil {
        fmt.Println(err.Error())
        return c.Render(400, "exercises_form", map[string]interface{}{
            "Error": []string{"Error, contact the admin"},
        })
    }
    if err := c.Render(200, "exercises_form", map[string]interface{}{"Exercises": exercises}); err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}
func (controller *ExercisesController) CreateExercise(c echo.Context) error {
    exercises, err := controller.ListExercisesService.ListExercises()
    name := c.FormValue("name")
    description := c.FormValue("description")
    link := c.FormValue("link")
    exercise := services.Exercise{Name: name, Description: description}
    messages, err := utils.Validate(exercise)
    if err != nil {
        return c.Render(400, "exercises_form", map[string]interface{}{
            "Errors": messages,
            "Exercises": exercises,
        })
    }
    messages, err = controller.CreateExerciseService.Create(name, description, link)
    if err != nil {
        fmt.Println(err.Error())
        return c.Render(400, "exercises_form", map[string]interface{}{
            "Errors": messages,
            "Exercises": exercises,
        })
    }
    exercises, err = controller.ListExercisesService.ListExercises()
    if err != nil {
        fmt.Println(err.Error())
        return c.Render(400, "exercises_form", map[string]interface{}{
            "Errors": []string{"Error, contact the admin"},
            "Exercises": exercises,
        })
    }
    return c.Render(200, "exercises_form", map[string]interface{}{
        "Messages": []string{"Exercise created successfully"},
        "Exercises": exercises,
    })
}
