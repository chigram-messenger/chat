package storage

import "fmt"

func (s *Storage) CreateMessage(idChat, idSender int64, message string) error {
	var op = "storage.message.CreateMessage"
	if _, err := s.db.Exec("INSERT INTO messages(id_chat, id_sender, message) values($1, $2, $3)", idChat, idSender, message); err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}
	return nil
}
