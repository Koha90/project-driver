package users

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int       `json:"id"`
	Username          string    `json:"username"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"createdAt"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pw)) == nil
}

func NewUser(username, password string) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Username:          username,
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
