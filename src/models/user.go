package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID          uint64    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Nick        string    `json:"nick,omitempty"`
	Email       string    `json:"email,omitempty"`
	Password    string    `json:"password,omitempty"`
	IsOwner     bool      `json:"isOwner,omitempty"`
	IsFollowing bool      `json:"isFollowing,omitempty"`
	Followers   uint64    `json:"followers,omitempty"`
	Following   uint64    `json:"following,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitzero"`
	Posts       []Post    `json:"posts,omitempty"`
}

func (u *User) Prepare(step string) error {
	if erro := u.format(step); erro != nil {
		return erro
	}

	if erro := u.validate(step); erro != nil {
		return erro
	}

	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("o nome é obrigatório")
	}

	if u.Nick == "" {
		return errors.New("o nick é obrigatório")
	}

	if u.Email == "" {
		return errors.New("o email é obrigatório")
	}

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("o email inserido é inválido")
	}

	if step == "signup" && u.Password == "" {
		return errors.New("a senha é obrigatório")
	}

	return nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if step != "signup" {
		return nil
	}

	passwordHashed, erro := security.Hash(u.Password)
	if erro != nil {
		return erro
	}
	u.Password = string(passwordHashed)

	return nil

}
