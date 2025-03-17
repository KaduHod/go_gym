package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"kaduhod/gym/app/controllers"
	"kaduhod/gym/app/services"
	"kaduhod/gym/app/utils"
	"kaduhod/gym/database"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
type Template struct {
    templates *template.Template
}
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func newTemplate() *Template {
    return &Template{
        templates: template.Must(template.New("").Funcs(template.FuncMap{
            "capitalize": utils.CapitalizeFirstLetter,
        }).ParseGlob("./views/*.tmpl")),
    }
}
type PortionMovementJoint struct {
    Muscle_group_name string
    Muscle_name string
    Muscle_id int
    Joint_name string
    Joint_id int
    Movement_name string
    Movement_id int
}
type MusclePortionViewData struct {
    Name string
    Itens []PortionMovementJoint
}
type MuscleGroupViewData struct {
    Name string
    Portions map[string]MusclePortionViewData
}
type SearchViewData struct {
    Groups map[string]MuscleGroupViewData
}
func main() {
    var db *sql.DB
    db = database.ConnetionMysql()
    e := echo.New()
    e.Renderer = newTemplate()
    listExercisesService := services.NewListExercisesService(db)
    createExerciseService := services.NewCreateExerciseService(db)
    signUpService := services.NewSignUpService(db)
    signInService := services.NewSignInService(db)
    buildExerciseService := services.NewBuildExerciseService(db)
    exerciseController := controllers.NewExerciseController(*listExercisesService, *createExerciseService, *buildExerciseService)
    userController := controllers.NewUserController(signUpService, signInService)
    baseController := controllers.NewController()
    currentTime := func() string {
		return time.Now().Format("02/01/2006 15:04:05.000") // Dia/mês/ano Hora:minuto:segundo.milissegundos
	}
    e.Static("/public", "public")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: fmt.Sprintf("[%s] path: ${uri} ${method} | status: ${status} | lat: ${latency}\n", currentTime()),
	}))
    e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
        TokenLookup: "header:X-CSRF-Token,form:csrf",
        CookieHTTPOnly: true,
    }))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3004"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
    e.Use(session.Middleware(sessions.NewCookieStore([]byte("shhiiiu"))))
    e.GET("/", baseController.Index)
    e.GET("/home", baseController.Home)
    e.POST("/muscles-joints-movements", func(c echo.Context) error {
        rows, err := db.Query(`SELECT
                mg.name AS muscle_group_name,
                mp.name AS muscle_name,
                mp.id AS muscle_id,
                a.name AS joint_name,
                a.id AS joint_id,
                m.name AS movement_name,
                m.id AS movement_id
            FROM articulation_movement_muscle amm
            INNER JOIN movements m ON m.id = amm.movement_id
            INNER JOIN articulations a ON a.id = amm.articulation_id
            INNER JOIN muscle_portion mp ON mp.id = amm.muscle_portion_id
            INNER JOIN muscle_group mg ON mp.muscle_group_id = mg.id
            ORDER BY mg.id`)
        if err != nil {
            return err
        }
        defer rows.Close()
        viewData := SearchViewData{
            Groups: make(map[string]MuscleGroupViewData),
        }
        for rows.Next() {
            var item PortionMovementJoint
            if err := rows.Scan(&item.Muscle_group_name, &item.Muscle_name, &item.Muscle_id, &item.Joint_name, &item.Joint_id, &item.Movement_name, &item.Movement_id); err == nil {
                if _, ok := viewData.Groups[item.Muscle_group_name]; !ok {
                    viewData.Groups[item.Muscle_group_name] = MuscleGroupViewData{
                        Name: item.Muscle_group_name,
                        Portions: make(map[string]MusclePortionViewData),
                    }
                }
                if _, ok := viewData.Groups[item.Muscle_group_name].Portions[item.Muscle_name]; !ok {
                    viewData.Groups[item.Muscle_group_name].Portions[item.Muscle_name] = MusclePortionViewData{
                        Name: item.Muscle_name,
                        Itens: make([]PortionMovementJoint, 0),
                    }
                }
                data := viewData.Groups[item.Muscle_group_name].Portions[item.Muscle_name]
                data.Itens = append(data.Itens, item)
                viewData.Groups[item.Muscle_group_name].Portions[item.Muscle_name] = data
            } else {
                fmt.Println(err)
                return c.Render(400, "search", map[string]interface{}{})
            }
        }
        if err := rows.Err(); err != nil {
            log.Println("Erro após iterar sobre rows:", err)
            return c.JSON(500, map[string]string{"error": "Erro ao iterar sobre dados"})
        }
        if err := c.Render(200, "search", viewData); err != nil {
            log.Println("Erro ao renderizar template:", err)
            return c.JSON(500, map[string]string{"error": "Erro ao renderizar página"})
        }
        return nil
    })
    e.POST("/exercises", exerciseController.ListExercises)
    e.POST("/exercises/details/:id", exerciseController.ExerciseDetail)
    e.GET("/signin", userController.SingInIndex)
    e.POST("/signin", userController.SingIn)
    e.GET("/signup", userController.SingUpIndex)
    e.POST("/signup", userController.SingUp)
    e.POST("/signout", baseController.LogOut)
    e.POST("/exercises/create", exerciseController.CreateExerciseForm)
    e.POST("/exercise", exerciseController.CreateExercise)
    e.POST("/exercise/:id", exerciseController.EditExerciseForm)
    e.PATCH("/exercise/:id/save", exerciseController.EditExercise)
    e.POST("/exercise/build/:id", exerciseController.Build)
    e.Logger.Fatal(e.Start(":3004"))
}
