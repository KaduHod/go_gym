package controllers

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Controller struct {}
func NewController() *Controller {
    return &Controller{}
}
func (c *Controller) getSession(e echo.Context) (*sessions.Session, error) {
    session, err := session.Get("session", e)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Erro ao pegar sessao")
    }
    return session, err
}
func (c *Controller) Index(e echo.Context) error {
    session, err := c.getSession(e)
    if err != nil {
        return err
    }
    userEmail := session.Values["user_email"]
    if userEmail == nil {
        return e.Render(200, "home", nil)
    }
    return e.Render(200, "home", map[string]interface{}{
        "UserEmail": userEmail,
    })
}
func (c *Controller) createSession(ctx echo.Context, user UserSignInDTO) (*sessions.Session, error) {
    session, err := session.Get("session", ctx)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Erro ao criar sessao")
        return session, err
    }
    session.Options = &sessions.Options{
        Path:     "/",
        MaxAge:   86400 * 7,
        HttpOnly: true,
    }
    session.Values["user_email"] = user.Email
    if err := session.Save(ctx.Request(), ctx.Response()); err != nil {
        return nil, err
    }
    return session, err
}
func (c *Controller) LogOut(ctx echo.Context) error {
    session, err := session.Get("session", ctx)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Erro ao destruir sessao")
        return err
    }
    session.Options = &sessions.Options{
        Path:     "/",
        MaxAge:   -1,
        HttpOnly: true,
    }
    if err := session.Save(ctx.Request(), ctx.Response()); err != nil {
        return err
    }
    return ctx.Redirect(303, "/")
}
