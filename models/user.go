package models

type User struct {
	Id        int32
	Username  string
	Password  string
	Age       int32
	CreatedAt int32
	UpdatedAt int32
}

func (u *User) InsertUser() (*User, error) {
	if err := Db.Create(u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (u *User) GetByUser() (*User, error) {
	user := User{}
	if err := Db.Where("username = ?", u.Username).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

func (u *User) ListUser() ([]User, error) {
	users := make([]User, 0)
	if err := Db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
