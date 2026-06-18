package storage

import (
	"chat/internal/domain"
	"chat/internal/service"
	"fmt"
)

func (s *Storage) CreateChat(idSender, idReceiver int64) error {
	var op = "storage.chat.CreateChat"
	if _, err := s.db.Exec("INSERT INTO chats(type, id_user_1, id_user_2) values($1, $2, $3)", service.PersonalChat, idSender, idReceiver); err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}
	return nil
}
func (s *Storage) GetChatId(idUser1, idUser2 int64) (int64, error) {
	var op = "storage.chat.GetChatId"
	var chatId int64
	row := s.db.QueryRow("SELECT id FROM chats WHERE (id_user_1=$1 AND id_user_2=$2) OR (id_user_1=$2 AND id_user_2=$1)", idUser1, idUser2)
	if err := row.Scan(&chatId); err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}
	return chatId, nil
}

func (s *Storage) GetAllChatsByID(id int64) ([]domain.Chat, error) {
	var op = "storage.chat.GetAllChatsByID"
	var chatsById = []domain.Chat{}
	var chatById = domain.Chat{}
	rows, err := s.db.Query("SELECT * FROM chats WHERE id_user_1=$1 OR id_user_2=$1", id)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}
	for rows.Next() {

		if err = rows.Scan(&chatById.Id, &chatById.Type, &chatById.IdUser1, &chatById.IdUser2); err != nil {
			return nil, fmt.Errorf("%s:%v", op, err)
		}

		chatsById = append(chatsById, chatById)
	}
	return chatsById, nil
}
