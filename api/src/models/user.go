package models

import (
	"errors"
	"strings"
	"time"
)

// User represents the user model/table
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare will validate and format an user model instance
func (u *User) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}

	u.format()
	return nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("Nome é obrigatório")
	}
	if u.Email == "" {
		return errors.New("Email é obrigatório")
	}
	if u.Nick == "" {
		return errors.New("Nick é obrigatório")
	}
	if u.Password == "" {
		return errors.New("Senha é obrigatório")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
