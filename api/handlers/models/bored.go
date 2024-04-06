package models

type BoredRes struct {
	Activity       string  `json:"response"`
}

type Post struct {
	Id          int64  `json:"id"`
	OwnerId     int64  `json:"owner_id"`
	Name        string `json:"name"`
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
}

type RegisterModel struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
