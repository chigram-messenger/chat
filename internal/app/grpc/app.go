package grpcapp

import (
	"chat/internal/config"
	"chat/internal/service"
	"chat/internal/transport"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	GRPCServer *grpc.Server
	port       int64
}

func New(log *slog.Logger, server config.GRPCServer, chatService service.ChatService, messageService *service.MessageService) *App {
	s := grpc.NewServer()
	transport.Register(s, chatService, messageService)
	return &App{
		GRPCServer: s,
		port:       server.Port,
		log:        log,
	}

}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	var op = "app.grpc.Run"
	addr := fmt.Sprintf(":%d", a.port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}
	a.log.Info("server starting...", slog.Int64("port", a.port))
	if err = a.GRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}
	return nil
}

func (a *App) Stop() {
	var op = "app.grpc.Stop"
	a.log.Info("server stopping...", slog.String("op", op))
	a.GRPCServer.GracefulStop()
}
