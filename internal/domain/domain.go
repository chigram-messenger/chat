package domain

type Chat struct {
	Id      int64
	Type    string
	IdUser1 int64
	IdUser2 int64
}
type ChatInfo struct {
	Id          int64
	IdCompanion int64
}
