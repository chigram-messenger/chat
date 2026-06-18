package service

import (
	"chat/internal/domain"
	"log/slog"
)

type Storage interface {
	CreateChat(idSender, idReceiver int64) error
	GetChatId(idUser1, idUser2 int64) (int64, error)
	CreateMessage(idChat, idSender int64, message string) error
	GetAllChatsByID(id int64) ([]domain.Chat, error)
}

type MessageService struct {
	log     *slog.Logger
	storage Storage
}
type ChatService struct {
	log     *slog.Logger
	storage Storage
}

func NewMessageService(log *slog.Logger, storage Storage) *MessageService {
	return &MessageService{log: log, storage: storage}
}
func NewChatService(log *slog.Logger, storage Storage) *ChatService {
	return &ChatService{log: log, storage: storage}
}
