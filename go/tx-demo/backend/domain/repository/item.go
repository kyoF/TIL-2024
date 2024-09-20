package repository

import "backend/domain/entity"

type Item interface {
	Get(itemId string) (*entity.Item, error)
	Insert(itemId, title, content string) error
	UpdateTitle(userId, title string) error
	UpdateContent(userId, content string) error
	Transaction
}
