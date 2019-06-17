package model

import "time"

type Article struct {
	ID              uint64    `db:"id"`
	Title           string    `db:"title"`
	Description     string    `db:"description"`
	Content         []byte    `db:"content"`
	MetaKeyWords    string    `db:"meta_keywords"`
	MetaDescription string    `db:"meta_description"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}