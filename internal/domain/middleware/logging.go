package middleware

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Log(logger *slog.Logger) gin.HandlerFunc {
	const op = "middleware.Log"
	log := logger.With(slog.String("op", op))

	log.Info("enabled")
	return func(c *gin.Context) {
		log.Info("[REQUEST]",
			slog.String("method", c.Request.Method),
			slog.String("endpoint", c.Request.URL.Path),
			slog.String("remote", c.Request.RemoteAddr),
		)

		defer func() {
			log.Info("[REQUEST] completed",
				slog.Int("code", c.Writer.Status()),
				slog.String("status", http.StatusText(c.Writer.Status())),
				slog.Bool("aborted", c.IsAborted()),
				slog.Int("body_size", c.Writer.Size()),
				slog.String("body", readN(c.Request.Body, 30)),
			)
		}()
		c.Next()
	}
}

func readN(r io.ReadCloser, max int) string {
	buf := make([]byte, max)

	n, err := r.Read(buf)
	if err != nil {
		return ""
	}

	return string(buf[:n]) + "..."
}
