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
}
func NewExerciseController(listExercisesService services.ListExercisesService, createExerciseService services.CreateExerciseService) *ExercisesController {
    return &ExercisesController{
        ListExercisesService: listExercisesService,
        CreateExerciseService: createExerciseService,
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
    viewData := ExerciseViewData{Exercises: exercises}
    fmt.Println(viewData)
    if err := c.Render(200, "list_exercises2", viewData); err != nil {
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
    mmjs, err := controller.ListExercisesService.ListExerciseJmm(exercise)
    if err != nil {
        return c.String(400, err.Error())
    }
    viewData := MMJViewData{
        Mmjs: mmjs,
        ImageURL: "public/images/exercises/"+strconv.Itoa(exercise.Id)+".jpeg",
        Exercise: exercise,
    }
    if err := c.Render(200, "exercise_detail", viewData); err != nil {
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
    exercise := services.Exercise{Name: name, Description: description}
    messages, err := utils.Validate(exercise)
    if err != nil {
        return c.Render(400, "exercises_form", map[string]interface{}{
            "Error": messages,
        })
    }
    messages, err = controller.CreateExerciseService.Create(name, description)
    if err != nil {
        return c.Render(400, "exercises_form", map[string]interface{}{
            "Error": messages,
        })
    }
    exercises, err = controller.ListExercisesService.ListExercises()
    if err != nil {
        fmt.Println(err.Error())
        return c.Render(400, "exercises_form", map[string]interface{}{
            "Error": []string{"Error, contact the admin"},
        })
    }
    return c.Render(200, "exercises_form", map[string]interface{}{
        "Message": []string{"Exercise created successfully"},
        "Exercises": exercises,
    })
}
