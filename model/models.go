package model

type Tokens struct {
	Id           uint
	UserId       uint
	AccessToken  string
	RefreshToken string
	Exp          int64
}

type User struct {
	Id    uint
	Login string
	Name  string
	Pass  string
}

type Task struct {
	Id          uint   `json:"id"`
	UserId      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed" sql:"DEFAULT:false; type:boolean; index" gorm:"not null"`
	StartDate   int64  `json:"start_date"`
	EndDate     int64  `json:"end_date"`
}
