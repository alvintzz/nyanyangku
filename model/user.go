package model

import (
	"fmt"
	"time"

	"github.com/alvintzz/nyanyangku/common/database"
	"github.com/alvintzz/nyanyangku/common/format"
)

//Container for all user related object.
type UserModel struct {
	Database database.Db
}

type User struct {
	ID       int    `db:"user_id"`
	Email    string `db:"user_email"`
	Password string `db:"user_password"`
	Status   int    `db:"status"`

	Name       string    `db:"user_name"`
	Phone      string    `db:"user_phone"`
	JoinTime   time.Time `db:"join_time"`
	UpdateTime time.Time `db:"update_time"`
}

// NewUserModel instantiates a new UserModel used for get/create user
func NewUserModel(db database.Db) *UserModel {
	return &UserModel{
		Database: db,
	}
}

const userTableDetail = `user_id, user_email, user_password, status, user_name, user_phone, join_time, update_time`
const userTableName = "tb_users"

const getUserByEmailQuery = "SELECT %s FROM %s WHERE user_email = $1"

// GetUserByEmail returns User Object from email inputted
func (m *UserModel) GetUserByEmail(email string) (User, error) {
	user := User{}
	query := fmt.Sprintf(getUserByEmailQuery, userTableDetail, userTableName)

	err := m.Database.DoAction().Get(&user, query, format.ToLower(email))
	if err == database.ErrNoRows {
		return user, nil
	} else if err != nil{
		return user, err
	}
	return user, nil
}
