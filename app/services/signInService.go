package services

import (
	"database/sql"
	"errors"
	"kaduhod/gym/app/utils"
	"strings"
)
type SignInService struct {
    db *sql.DB
}
func NewSignInService(db *sql.DB) *SignInService {
    return &SignInService{
        db: db,
    }
}
func (this *SignInService) SignIn(email string, password string) ([]string, error) {
    errorMessages, err := this.validate(email, password)
    if len(errorMessages) > 0 {
        return errorMessages, errors.New("Validation errors")
    }
    if err != nil {
        return nil, err
    }
    var user User
    err = this.db.QueryRow("SELECT id, name, email, password FROM user WHERE email = ?", strings.ToLower(email)).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return []string{"User not found"}, errors.New("User not found")
        }
        return []string{"Error checking user"}, err
    }
    if !CheckPasswordHash(password, user.Password) {
        return []string{"Invalid password"}, errors.New("Invalid password")
    }
    return nil, nil
}
type UserSignInDTO struct {
    Email string `form:"email" json:"email" validate:"required,email"`
    Password string `form:"password" json:"password" validate:"required,min=8"`
}
func (this *SignInService) validate(email string, password string) ([]string, error) {
    errorMessages, err := utils.Validate(UserSignInDTO{Email: email, Password: password})
    if err != nil {
        return errorMessages, err
    }
    if len(errorMessages) > 0 {
        return errorMessages, errors.New("Validation errors")
    }
    return errorMessages, nil
}
