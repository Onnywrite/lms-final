package middleware

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recover(logger *slog.Logger) gin.HandlerFunc {
	const op = "middleware.Recover"
	log := logger.With(slog.String("op", op))

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)

				headerJson, _ := json.Marshal(&c.Request.Header)

				body, _ := io.ReadAll(c.Request.Body)

				log.Error("recovered a panic",
					slog.Any("err", err),
					slog.String("header", string(headerJson)),
					slog.String("body", string(body)),
				)
			}
		}()

		c.Next()
	}
}
