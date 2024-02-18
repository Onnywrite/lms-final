package middleware

import (
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORS(logger *slog.Logger, allowOrigin []string) gin.HandlerFunc {
	const op = "middleware.CORS"
	log := logger.With(slog.String("op", op))

	origins := strings.Join(allowOrigin, ", ")
	log.Info("enabled", slog.String("allowed_origins", origins))
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", origins)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("Access-Control-Max-Age", "86400")

		c.Next()
	}
}

func CORSAllOrigins(logger *slog.Logger) gin.HandlerFunc {
	return CORS(logger, []string{"*"})
}
