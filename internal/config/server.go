package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppServer struct {
	server *http.Server
}

func NewServer(routers *gin.Engine) *AppServer {
	return &AppServer{
		server: &http.Server{
			Addr:    envVars.Port,
			Handler: routers,
		},
	}
}

func (as *AppServer) StartServer() error {
	if err := as.server.ListenAndServe(); err != nil {
		return err
	}
	Logger.Info("Server started on port: %s", envVars.Port)
	return nil
}
