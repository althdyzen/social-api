package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Body       string    `json:"body,omitempty"`
	AuthorId   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitzero"`
	AuthorName string    `json:"authorName,omitempty"`
	IsAuthor   bool      `json:"isAuthor,omitempty"`
}

func (p *Post) Prepare() error {
	p.format()

	if erro := p.validate(); erro != nil {
		return erro
	}

	return nil
}

func (p *Post) validate() error {
	if p.Title == "" {
		return errors.New("o título não pode estar em branco")
	}
	if p.Body == "" {
		return errors.New("o conteudo não pode estar em branco")
	}

	return nil
}

func (p *Post) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Body = strings.TrimSpace(p.Body)
}
