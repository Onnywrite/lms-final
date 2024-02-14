package handlers

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func LogMiddleware(log *slog.Logger) func(*gin.Context) {
	log.Info("enabled logMiddleware")

	return func(c *gin.Context) {
		log = log.With(slog.String("op", "logMiddleware"))
		defer func() {
			log.Info("request completed",
				slog.String("method", c.Request.Method),
				slog.String("endpoint", c.Request.URL.Path),
				slog.String("remote", c.Request.RemoteAddr),
			)
		}()

		c.Next()
	}
}
