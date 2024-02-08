package restful

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type FrontendServer struct {
	log  *slog.Logger
	mux  *gin.Engine
	port int
}

const (
	path          = "./resources/"
	indexHTML     = "index.html"
	indexHTMLPath = path + indexHTML
	scriptJS      = "script.js"
	scriptJSPath  = path + scriptJS
)

func NewFrontend(logger *slog.Logger, port int) *FrontendServer {
	mux := gin.Default()
	mux.StaticFile(scriptJS, scriptJSPath)
	mux.GET("/", func(c *gin.Context) {
		c.HTML(200, indexHTML, nil)
	})
	mux.GET("/status/", handleStatusHTML)
	mux.GET("/powers/", handlePowersHTML)
	logger.Debug("New restful.FrontendServer is ready to give HTML-pages")

	return &FrontendServer{
		log:  logger,
		port: port,
		mux:  mux,
	}
}

func (f *FrontendServer) Start() {
	f.mux.LoadHTMLFiles(indexHTMLPath)
	go f.mux.Run(fmt.Sprintf(":%d", f.port))
	f.log.Info("restful.FrontendServer started listening and serving")
}

func handleStatusHTML(c *gin.Context) {

}
func handlePowersHTML(c *gin.Context) {

}
