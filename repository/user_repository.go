package repository

import (
	// "fmt"

	"fmt"

	"gorm.io/gorm"
)

type userRepositoryDb struct {
	db *gorm.DB
}

func NewUserRepositoryDb(db *gorm.DB) userRepositoryDb {
	return userRepositoryDb{db: db}
}

// CreateUser implements UserRepository.
func (u userRepositoryDb) CreateUser(data User_info) (string, error) {
	fmt.Println("Data for create : ", data)
	result := u.db.Exec("INSERT INTO user_info (user_id, username, created_at, lasted_login) VALUES(?, ?, ?, ?);", data.User_id, data.Username, data.Created_at, data.Lasted_login)

	fmt.Println("result :", result)
	fmt.Println("Error : ", result.Error)
	if result.Error != nil {
		fmt.Println("Error : ", result.Error)
		return "Error", result.Error
	}
	return "Text", nil
}

// GetById implements UserRepository.
func (u userRepositoryDb) GetById(id string) (User_info, error) {
	data := User_info{}

	result := u.db.Raw("SELECT * from user_info WHERE user_id = ?", id).Scan(&data)

	if result.Error != nil {
		fmt.Println("Error : ", result.Error)
		return data, result.Error
	}
	return data, nil
}

// UpdateLastLogin implements UserRepository.
func (u userRepositoryDb) UpdateLastLogin(id string) error {	
	result := u.db.Exec("UPDATE user_info SET lasted_login = datetime('now', 'localtime') WHERE user_id=?", id)

	if result.Error != nil {
		fmt.Println("Error : ", result.Error)
		return result.Error
	} else if result.RowsAffected <= 0 {
		err := fmt.Errorf("don't have this record")
		return err
	}
	return nil
}

func (u userRepositoryDb) DeleteUser(id string) error {
	result := u.db.Exec("DELETE FROM user_info WHERE user_id=? ;", id)
	fmt.Println("deleted row record : ", result.RowsAffected)

	if result.Error != nil {
		fmt.Println("Error : ", result.Error)
		return result.Error
	} else if result.RowsAffected <= 0 {
		err := fmt.Errorf("don't have this record")
		return err
	}
	return nil
}
