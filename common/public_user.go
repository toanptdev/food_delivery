package common

type User struct {
	SqlModel
	Role string `json:"role" gorm:"column:role"`
}

func (u User) TableName() string {
	return "users"
}
