package repository

type User_info struct {
	// user_id	username	created_at	lasted_login
	User_id      string `gorm:"primary_key user_id"`
	Username     string `db:"username"`
	Created_at   string `db:"created_at"`
	Lasted_login string `db:"lasted_login"`
}

// TableName overrides the table name used by User to `user_info`
func (User_info) TableName() string {
	return "user_info"
}

type UserRepository interface {
	GetById(string) (User_info, error)
	CreateUser(User_info) (string, error)
	UpdateLastLogin(string) error
	DeleteUser(string) error
}
