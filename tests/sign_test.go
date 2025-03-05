package tests

import (
	"fmt"
	"kaduhod/gym/app/services"
	"kaduhod/gym/database"
	"testing"
)
var signUpService *services.SignUpService
var signInService *services.SignInService

func TestSignUpError(t *testing.T) {
    db := database.ConnetionMysql()
    signUpService = services.NewSignUpService(db)
    messages, err := signUpService.SignUp("test", "test@mail.com", "test")
    if err == nil || len(messages) == 0 {
        fmt.Println("Deveria ter retornado erro", err)
        t.Fail()
        return
    }
}
func TestSignUpSuccess(t *testing.T) {
    db := database.ConnetionMysql()
    signUpService = services.NewSignUpService(db)
    messages, err := signUpService.SignUp("test", "test@mail2.com", "123456789")
    if err != nil || len(messages) > 0 {
        fmt.Println("Error signing up: ", err)
        fmt.Println("Error messages: ", messages)
        t.Fail()
        return
    }
}
func TestSignInError(t *testing.T) {
    db := database.ConnetionMysql()
    signInService = services.NewSignInService(db)
    _, err := signInService.SignIn("test", "test")
    if err == nil {
        fmt.Println("Deveria ter retornado erro", err)
        t.Fail()
        return
    }
}
func TestSignInSuccess(t *testing.T) {
    db := database.ConnetionMysql()
    signInService = services.NewSignInService(db)
    _, err := signInService.SignIn("teste@mail.com", "12345678")
    if err != nil  {
        fmt.Println("Error signing in: ", err)
        t.Fail()
        return
    }
}
