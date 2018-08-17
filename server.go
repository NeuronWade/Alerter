package main

import (
	"alerter/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
	"alerter/obtime"
)

const(
	LoopInterval = 30 * time.Second
	ClaneInterval = 1 * time.Hour
)

type Server struct {
	echo   *echo.Echo
	ListenAddr string

	loopTicker  *time.Ticker
	cleanticker *time.Ticker
}

func New() *Server {
	return &Server{
		ListenAddr:  config.App.ServerAddr,
		loopTicker:  time.NewTicker(LoopInterval),
		cleanticker: time.NewTicker(ClaneInterval),
	}
}

func (s *Server) OnInit() {
	s.echo = echo.New()
	s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: log.Output(),
	}))

	s.echo.POST("/api/statics_alert", StaticsAlert)

	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	log.Debugf("ops listen begin\n")
}

func (s *Server) OnDestroy() {
	s.echo.Close()
}

func (s *Server) Run() {
	s.echo.Start(s.ListenAddr)
	for {
		select{
		case <- s.loopTicker.C:
			//s.loop()
			case <- s.cleanticker.C:
			s.clean()
		}
	}
}

// NOTE: 查找创建时间与更新时间相隔一定时间的Stack信息
func (s *Server) clean() {
	for stackHead, stackInfo := range serverStack.StackList {
		if stackInfo.SendFlag && (stackInfo.UpdateTime - obtime.Now().Unix()) > 3600 {
			delete(serverStack.StackList, stackHead)
		}
	}
}