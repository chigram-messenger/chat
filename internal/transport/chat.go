package transport

import (
	"context"
	"fmt"

	"github.com/chigram-messenger/protos/gen/pb/chat/chat"
)

func (s *ChatServer) GetAllByID(ctx context.Context, r *chat.GetAllByIDRequest) (*chat.GetAllByIDResponse, error) {
	var chatResp *chat.ChatInfo

	var chatsResp []*chat.ChatInfo
	id := r.GetUserId()
	chatsInfo, err := s.serv.GetAllByID(id)
	if err != nil {
		return &chat.GetAllByIDResponse{}, fmt.Errorf("Internal error")
	}
	for _, chatInfo := range chatsInfo {
		chatResp = &chat.ChatInfo{ChatId: chatInfo.Id, CompanionId: chatInfo.IdCompanion}
		chatsResp = append(chatsResp, chatResp)
	}

	return &chat.GetAllByIDResponse{Chats: chatsResp}, nil
}

//func (s *ChatServer) Get(r *chat.GetRequest) (*chat.GetResponse, error) {
//	users_id := r.GetUserId()
//	s.serv.GetMessagesByChatId()
//}
