package usecase

import (
	"backend/application/dto"
	"backend/domain/repository"
)

type Item interface {
	Get(itemId string) (dto.Item, error)
	Insert(itemId, title, content string) error
	UpdateTitle(itemId, title string) error
	UpdateContent(itemId, content string) error
}

type item struct {
	itemRepository repository.Item
}

func NewItem(itemRepository repository.Item) Item {
	return &item{itemRepository: itemRepository}
}

func (i *item) Get(itemId string) (dto.Item, error) {
	item, err := i.itemRepository.Get(itemId)
	if err != nil {
		return dto.Item{}, err
	}

	return dto.Item{
		ItemId:  item.ItemId,
		Title:   item.Title,
		Content: item.Content,
	}, nil
}

func (i *item) Insert(userId, title, content string) error {
	return i.itemRepository.Transaction(func() error {
		return i.itemRepository.Insert(userId, title, content)
	})
}

func (i *item) UpdateTitle(userId, tilte string) error {
	return i.itemRepository.Transaction(func() error {
		return i.itemRepository.UpdateTitle(userId, tilte)
	})
}

func (i *item) UpdateContent(userId, content string) error {
	return i.itemRepository.Transaction(func() error {
		return i.itemRepository.UpdateContent(userId, content)
	})
}
