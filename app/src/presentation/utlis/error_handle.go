package utils

import (
	commons "app/domain/commons/messages"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleError(ctx echo.Context, err error) {
	var customErr *commons.CustomError
	if errors.As(err, &customErr) {
		switch customErr.Type {
		case commons.TypeValidation:
			ctx.JSON(http.StatusBadRequest, map[string]string{"error": customErr.Error()})
		case commons.TypeNotFound:
			ctx.JSON(http.StatusNotFound, map[string]string{"error": customErr.Error()})
		case commons.TypeConflict:
			ctx.JSON(http.StatusConflict, map[string]string{"error": customErr.Error()})
		case commons.TypeUnAuthorized:
			ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		case commons.TypeForbidden:
			ctx.JSON(http.StatusForbidden, map[string]string{"error": customErr.Error()})
		case commons.TypeBadRequest:
			ctx.JSON(http.StatusBadRequest, map[string]string{"error": customErr.Error()})
		case commons.TypeInternal:
			ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		}

		return
	}

	ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
}
