package service

import "log/slog"

func (s *MessageService) Create(isChat bool, idChat, idSender, idReceiver int64, message string) bool {
	var op = "service.Message.Create"
	if !isChat {
		if err := s.storage.CreateChat(idSender, idReceiver); err != nil {
			s.log.Error("dont created new chat", slog.String("op", op), slog.Any("error", err))
			return false
		}
		id, err := s.storage.GetChatId(idSender, idReceiver)
		if err != nil {
			s.log.Error("dont got id chat", slog.String("op", op), slog.Any("error", err))
			return false
		}
		idChat = id
	}

	if err := s.storage.CreateMessage(idChat, idSender, message); err != nil {
		s.log.Error("dont created message", slog.String("op", op), slog.Any("error", err))
		return false
	}
	return true

}
