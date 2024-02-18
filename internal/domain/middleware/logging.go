package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func Log(log *slog.Logger) gin.HandlerFunc {
	log.Info("enabled logMiddleware")

	return func(c *gin.Context) {
		log.Info("request completed",
			slog.String("method", c.Request.Method),
			slog.String("endpoint", c.Request.URL.Path),
			slog.String("remote", c.Request.RemoteAddr),
		)

		c.Next()
	}
}
