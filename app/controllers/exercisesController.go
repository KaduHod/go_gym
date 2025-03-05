package controllers

import (
	"fmt"
	"kaduhod/gym/app/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ExercisesController struct {
    Controller
    ListExercisesService services.ListExercisesService
}

func NewExerciseController(listExercisesService services.ListExercisesService) *ExercisesController {
    return &ExercisesController{
        ListExercisesService: listExercisesService,
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
