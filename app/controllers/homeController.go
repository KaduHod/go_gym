package controllers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)
type HomeController struct {
    Controller
}
func NewHomeController() *HomeController {
    return &HomeController{}
}
type HomeData struct {
    Messages []string
    Data map[string]interface {}
}
func (c *HomeController) HomeIndex(e echo.Context) error {
    if err := e.Render(200, "home", nil); err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}
