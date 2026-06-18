package transport

import (
	"chat/internal/service"

	"github.com/chigram-messenger/protos/gen/pb/chat/chat"
	"github.com/chigram-messenger/protos/gen/pb/chat/message"
	"google.golang.org/grpc"
)

type MessageService interface {
	Create(isChat bool, idChat, idSender, idReceiver int64, message string) bool
}
type MessageServer struct {
	message.UnimplementedMessageServer
	serv MessageService
}
type ChatServer struct {
	chat.UnimplementedChatServer
	serv service.ChatService
}

func Register(s *grpc.Server, chatSrv service.ChatService, messageSrv MessageService) {
	message.RegisterMessageServer(s, &MessageServer{serv: messageSrv})
	chat.RegisterChatServer(s, &ChatServer{serv: chatSrv})

}
