package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](i T) ([]string, error) {
    var errorsMessages []string
    validate := validator.New(validator.WithRequiredStructEnabled())
    var err error
    err = validate.Struct(i)
    if err != nil {
        var invalidErrors validator.ValidationErrors
        if errors.As(err, &invalidErrors) {
            for _, invalidError := range invalidErrors {
                if invalidError.Tag() == "required" {
                    errorsMessages = append(errorsMessages, "Field "+invalidError.Field()+" is required")
                }
                if invalidError.Tag() == "min" {
                    errorsMessages = append(errorsMessages, "Field "+invalidError.Field()+" must be at least "+invalidError.Param())
                }
                if invalidError.Tag() == "email" {
                    errorsMessages = append(errorsMessages, "Field "+invalidError.Field()+" must be a valid email")
                }
                if invalidError.Tag() == "eqfield" {
                    errorsMessages = append(errorsMessages, "Field "+invalidError.Field()+" must be equal to "+invalidError.Param())
                }
            }
        }
        return errorsMessages, err
    }
    return errorsMessages, nil
}
