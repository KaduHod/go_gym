package services

import (
	"database/sql"
	"errors"
	"kaduhod/gym/app/utils"
	"strings"
)

type SignUpService struct {
    db *sql.DB
}

func NewSignUpService(db *sql.DB) *SignUpService {
    return &SignUpService{
        db: db,
    }
}
type User struct {
    Id       int
    Name     string `validate:"required,ne=''"`
    Email    string `validate:"required,email"`
    Password string `validate:"required,min=8"`
}
func (this *SignUpService) SignUp(name string, email string, password string) ([]string, error) {
    errorMessages, err := this.validate(name, email, password)
    if err != nil {
        return errorMessages, err
    }
    exists, err := this.checkEmail(email)
    if err != nil {
        return []string{"Error checking email"}, err
    }
    if exists {
        return []string{"Email unavailable"}, errors.New("Email already exists")
    }
    return []string{}, this.save(name, email, password)
}
func (this *SignUpService) validate(name string, email string, password string) ([]string, error) {
    errorMessages, err := utils.Validate(User{Name: name, Email: email, Password: password})
    if err != nil {
        return errorMessages, err
    }
    if len(errorMessages) > 0 {
        return errorMessages, errors.New("Validation errors")
    }
    return errorMessages, nil
}
func (this *SignUpService) save(name string, email string, password string) error {
    hashedPassword, err := HashPassword(password)
    if err != nil {
        return err
    }
    _, err = this.db.Exec("INSERT INTO user (name, email, password) VALUES (?, ?, ?)", name, strings.ToLower(email), hashedPassword)
    return err
}
func (this *SignUpService) checkEmail(email string) (bool, error) {
    var exists bool
    err := this.db.QueryRow("SELECT EXISTS(SELECT * FROM user WHERE email = ?)", email).Scan(&exists)
    return exists, err
}
