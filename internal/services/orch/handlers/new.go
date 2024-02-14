package handlers

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"github.com/Onnywrite/lms-final/internal/domain/models"
	"github.com/gin-gonic/gin"
)

var (
	invalidTimingsErr = errors.New("invalid timings! They must be > 0")
	invalidSymbolsErr = errors.New("expression contains restricted characters")
)

type ExpressionSaver interface {
	Save(*models.ParsedExpression) (models.ProcessedExpression, error)
}

func PostNew(saver ExpressionSaver) func(*gin.Context) {
	return func(c *gin.Context) {
		var expr models.Expression
		if err := c.BindJSON(&expr); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &models.JsonErr{
				Error: err.Error(),
				Msg:   "invalid request format or fields",
			})
			return
		}

	}
}

func parseExpression(e *models.Expression) (*models.ParsedExpression, error) {
	if err := validateExpression(e); err != nil {
		return nil, err
	}

	return nil, nil
}

func validateExpression(e *models.Expression) error {
	if e.DivisionTime <= 0 || e.MultiplicationTime <= 0 ||
		e.AdditionTime <= 0 || e.SubtractionTime <= 0 {
		return invalidTimingsErr
	}

	e.Expression = removeSpaces(e.Expression)
	if !regexp.MustCompile(`[0-9()+\-*/]+`).MatchString(e.Expression) {
		return invalidSymbolsErr
	}

	return nil
}

func removeSpaces(expr string) string {
	sb := strings.Builder{}
	sb.Grow(len(expr))
	for _, r := range expr {
		if !unicode.IsSpace(r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
