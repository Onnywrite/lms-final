package handlers

import (
	"errors"
	"github.com/Onnywrite/lms-final/internal/domain/models"
	"github.com/Onnywrite/lms-final/internal/services/orch/rpn"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	invalidTimingsErr = errors.New("invalid timings! They must be > 0")
)

type ExpressionSaver interface {
	Save( /**models.ParsedExpression*/ ) (models.ProcessedExpression, error)
}

func PostNew(saver ExpressionSaver) func(*gin.Context) {
	return func(c *gin.Context) {
		var err error
		var expr models.Expression
		if err = c.BindJSON(&expr); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &models.JsonErr{
				Error: err.Error(),
				Msg:   "invalid request format or fields",
			})
			return
		}

		if err = validateSettings(&expr); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &models.JsonErr{
				Error: err.Error(),
				Msg:   "invalid request fields. Timings must be greater than zero",
			})
		}

		// Temporary
		var rpnExpr string
		if rpnExpr, err = rpn.FromInfix(expr.Expression); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &models.JsonErr{
				Error: err.Error(),
				Msg:   "",
			})
		}

		// Temporary
		c.JSON(http.StatusOK, gin.H{
			"rpn_expression": rpnExpr,
		})
	}
}

func validateSettings(e *models.Expression) error {
	if e.DivisionTime <= 0 || e.MultiplicationTime <= 0 ||
		e.AdditionTime <= 0 || e.SubtractionTime <= 0 {
		return invalidTimingsErr
	}
	return nil
}
