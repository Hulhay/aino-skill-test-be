package model

type User struct {
	ID       int64  `json:"id" gorm:"type:int primary key auto_increment"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	IsLogin  bool   `json:"is_login"`
	Password string
}
