package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/corentindeboisset/neutrino/pkg/logger"
	"github.com/corentindeboisset/neutrino/pkg/server/apiv1"
	"github.com/corentindeboisset/neutrino/pkg/transfermgr"
	"github.com/corentindeboisset/neutrino/web"
	"github.com/gin-gonic/gin"
)

func StartServer(serverConfig *Config, debugMode bool) error {
	if debugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(clientMgmtMiddleware())
	router.Use(transfermgr.TransferMgmtMiddleware())

	if logger.GetLogLevel() >= 1 {
		var logConfig gin.LoggerConfig
		logConfig.Formatter = func(param gin.LogFormatterParams) string {
			return fmt.Sprintf(
				"%s - [%s] %s %s %s %d %s \"%s\" %s\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}
		logConfig.Output = log.Writer()
		router.Use(gin.LoggerWithConfig(logConfig))
	}

	router.StaticFS("/static", http.FS(web.StaticFiles))

	// ./ automatically serves ./index.html
	router.StaticFileFS("/", "./", http.FS(web.RootFiles))
	router.StaticFileFS("/favicon_plain.png", "./favicon_plain.png", http.FS(web.RootFiles))

	// Setup the apiv1 group of routes
	apiv1.DecorateRouter(router)
	// Next versions of the API should be added here

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port),
		Handler: router,
	}

	logger.InfoLog("Starting the server on http://%s:%d", serverConfig.Host, serverConfig.Port)
	err := server.ListenAndServe()
	if err != nil {
		logger.ErrorLog("The server closed with the following status: %v", err)
	}

	return nil
}
