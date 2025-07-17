package entities

type UsersRegisterReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersRegisterRes struct {
	Id       uint64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}
