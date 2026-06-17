package transport

import (
	"context"
	"fmt"

	"github.com/chigram-messenger/protos/gen/pb/chat/message"
)

func (s *MessageServer) Create(ctx context.Context, r *message.CreateRequest) (*message.CreateResponse, error) {
	if ok := s.serv.Create(r.IsChat, r.IdChat, r.IdSender, r.IdReceiver, r.Message); !ok {
		return &message.CreateResponse{Status: false}, fmt.Errorf("status error")
	}
	return &message.CreateResponse{Status: true}, nil
}
