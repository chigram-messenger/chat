package storage

import "fmt"

func (s *Storage) CreateChat(idSender, idReceiver int64) error {
	var op = "storage.chat.CreateChat"
	if _, err := s.db.Exec("INSERT INTO chats(id_user_1, id_user_2) values($1, $2)", idSender, idReceiver); err != nil {
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
