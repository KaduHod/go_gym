package controllers

import (
	"fmt"
	"kaduhod/gym/app/services"

	"github.com/labstack/echo/v4"
)

type UserController struct {
    Controller
    signUpService *services.SignUpService
    signInService *services.SignInService
}
func NewUserController(signUpService *services.SignUpService, signInService *services.SignInService) *UserController {
    return &UserController{
        signUpService: signUpService,
        signInService: signInService,
    }
}
type UserSignInDTO struct {
    Email string `form:"email" json:"email" validate:"required,email"`
    Password string `form:"password" json:"password" validate:"required,min=8"`
}
func (c *UserController) SingInIndex(e echo.Context) error {
    if err := e.Render(200, "signin", nil); err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}
type SignInData struct {
    Errors []string
    UserDTO UserSignInDTO
    Messages []string
}
func (c *UserController) SingIn(e echo.Context) error {
    var user UserSignInDTO
    if err := e.Bind(&user); err != nil {
        fmt.Println(err.Error())
        return err
    }
    errors, err := c.signInService.SignIn(user.Email, user.Password)
    if err != nil {
         errorsObj := SignInData{Errors: errors, UserDTO: user}
        if err := e.Render(400, "signin", errorsObj); err != nil {
            fmt.Println(err.Error())
            return err
        }
        return err
    }
    _, err = c.createSession(e, user);
    if err != nil {
        fmt.Println(err)
        return err
    }
    //viewData := HomeData{Messages: []string{"You are now logged in!"}}
    e.Redirect(303,"/")
    return nil
}
type UserDTO struct {
    Name string `form:"name" json:"name" validate:"required,min=5,ne=''"`
    Email string `form:"email" json:"email" validate:"required,email"`
    Password string `form:"password" json:"password" validate:"required,min=8"`
    ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" validate:"required,min=8,eqfield=Password"`
}
func (c *UserController) SingUpIndex(e echo.Context) error {
    if err := e.Render(200, "signup", nil); err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}
type SignUpData struct {
    Errors []string
    UserDTO UserDTO
}
func (c *UserController) SingUp(e echo.Context) error {
    var user UserDTO
    if err := e.Bind(&user); err != nil {
        fmt.Println(err.Error())
        return err
    }
    errorMessages, err := c.signUpService.SignUp(user.Name, user.Email, user.Password)
    if err != nil {
        errorsObj := SignUpData{Errors: errorMessages, UserDTO: user}
        if err := e.Render(400, "signup", errorsObj); err != nil {
            fmt.Println(err.Error())
            return err
        }
        return err
    }
    SignInData := SignInData{
        UserDTO: UserSignInDTO{
            Email: user.Email,
            Password: user.Password,
        },
        Messages: []string{"Your account has been created successfully. You can now sign in!"},
    }
    e.Render(200, "signin", SignInData)
    return nil
}
