package common

import (
	"fmt"
	"github.com/cesc1802/auth-service/pkg/i18n"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode      int                    `json:"status_code"`
	RootErr         error                  `json:"-"`
	Message         string                 `json:"message"`
	Log             string                 `json:"log"`
	Key             string                 `json:"error_key"`
	ValidationError []ValidationErrorField `json:"ve,omitempty"`
}

type ValidationErrorField struct {
	Field        string `json:"field,omitempty"`
	Tag          string `json:"tag,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

var ErrRecordNotFound = NewCustomError(nil, "record not found", "ERR_RECORD_NOT_FOUND")

func ValidationError(msg string, key string, ve []ValidationErrorField) *AppError {
	appErr := NewErrorResponse(nil, msg, msg, key)
	appErr.ValidationError = ve
	return appErr
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ERR_INVALID_REQUEST")
}

func ErrInternal(err error) *AppError {
	return NewErrorResponse(err, "internal error", err.Error(), "ERR_INTERNAL")
}

//func ErrRecordNotFound() *AppError {
//	return NewCustomError(nil, "record not found", "ERR_RECORD_NOT_FOUND")
//}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_LIST_%s", strings.ToUpper(entity)),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_DELETE_%s", strings.ToUpper(entity)),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_GET_%s", strings.ToUpper(entity)),
	)
}

func ErrEntityExisting(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("ERR_%s_ALREADY_EXIST", strings.ToUpper(entity)),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_CREATE_%s", strings.ToUpper(entity)),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Update %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_UPDATE_%s", strings.ToUpper(entity)),
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("You have no permission"),
		fmt.Sprintf("ERR_NO_PERMISSION"),
	)
}

func translateToAppVE(i18n *i18n.AppI18n, lang string,
	valErrors validator.ValidationErrors, errCode string) []ValidationErrorField {

	res := make([]ValidationErrorField, len(valErrors))
	for i, valErr := range valErrors {
		res[i] = ValidationErrorField{
			Field:        valErr.Field(),
			Tag:          valErr.Tag(),
			ErrorMessage: i18n.MustLocalize(lang, fmt.Sprintf("%v.%v", errCode, valErr.Tag()), nil),
		}
	}
	return res
}

func HandleValidationErrors(language string, i18n *i18n.AppI18n, valErrors validator.ValidationErrors) *AppError {
	appErr := ValidationError(
		i18n.MustLocalize(language, "COM0005", nil),
		"ERR_VALIDATION_REQUEST",
		translateToAppVE(i18n, language, valErrors, "COM0005"),
	)
	return appErr
}

func HandleAppError(language string, i18n *i18n.AppI18n, err error) *AppError {
	appErr := err.(*AppError)
	appErr.Message = i18n.MustLocalize(language, appErr.Key, nil)
	return appErr

}

func MustError(err error) {
	if err != nil {
		panic(err)
	}
}
