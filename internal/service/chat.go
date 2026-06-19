package service

import (
	"chat/internal/domain"
	"log/slog"
)

const (
	PersonalChat = "personal"
	GroupChat    = "group"
)

func (s *ChatService) GetAllByID(id int64) ([]domain.ChatInfo, error) {
	var op = "service.chat.GetAllByID"
	var idCompanion int64
	chatInfo := domain.ChatInfo{}
	chatsInfo := make([]domain.ChatInfo, 0)
	chats, err := s.storage.GetAllChatsByID(id)
	if err != nil {
		s.log.Error("dont get chats for user", slog.String("op", op), slog.Any("error", err))
		return nil, err
	}
	for _, chat := range chats {
		switch chat.Type {
		case PersonalChat:
			if chat.IdUser1 == id {
				idCompanion = chat.IdUser2
			} else {
				idCompanion = chat.IdUser1
			}
		default:
			s.log.Error("i dont know this type chat", slog.String("op", op))
			idCompanion = 0
		}
		chatInfo = domain.ChatInfo{Id: id, IdCompanion: idCompanion}
		chatsInfo = append(chatsInfo, chatInfo)
	}
	return chatsInfo, nil

}

func (s *ChatService) GetMessagesByChatId(usersId []int64) {
	
}
