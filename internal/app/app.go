package app

import (
	grpcapp "chat/internal/app/grpc"
	"chat/internal/config"
	"chat/internal/service"
	"chat/internal/storage"
	"chat/internal/storage/postgres"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, server config.GRPCServer, p config.Postgres) *App {

	params := postgres.Params{
		Host:     p.Host,
		Port:     p.Port,
		User:     p.User,
		DBName:   p.DBName,
		Password: p.Password,
		Sslmode:  p.Sslmode,
	}
	db, err := postgres.Conn(params)
	if err != nil {
		panic("dot connection postgres database, err:" + err.Error())
	}
	storage := storage.New(db)

	chatService := service.NewChatService(log, storage)
	messageService := service.NewMessageService(log, storage)
	grpcApp := grpcapp.New(log, server, chatService, messageService)
	return &App{GRPCServer: grpcApp}
}
