package http

import (
	"strconv"

	"github.com/ali-shokoohi/notes/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server struct {
	cfg    *config.Config
	engine *gin.Engine
}

type Server interface {
	Launch() error
}

func NewServer(
	cfg *config.Config,
	engine *gin.Engine,
) Server {
	return &server{
		cfg:    cfg,
		engine: engine,
	}
}

func (s *server) Launch() error {
	listenAddress := ":" + strconv.Itoa(s.cfg.Port)
	if err := s.engine.Run(listenAddress); err != nil {
		s.cfg.Logger.Error("failed at starting of Gin:", zap.Error(err))
		return err
	}
	return nil
}
