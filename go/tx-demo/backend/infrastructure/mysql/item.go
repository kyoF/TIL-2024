package mysqlinfra

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/infrastructure/mysql/models"
)

type item struct {
	*dbClient
}

func NewItem(dbClient *dbClient) repository.Item {
	return &item{
		dbClient: dbClient,
	}
}

func (i *item) Get(itemId string) (*entity.Item, error) {
	var item models.Item
	err := i.dbClient.DB().Where("item_id = ?", itemId).First(&item).Error
	if err != nil {
		return nil, err
	}

	return &entity.Item{
		ItemId:  item.ItemID,
		Title:   item.Title,
		Content: item.Content,
	}, nil
}

func (i *item) Insert(itemId, title, content string) error {
	item := &models.Item{
		ItemID:  itemId,
		Title:   title,
		Content: content,
	}

	err := i.dbClient.DB().Create(item).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *item) UpdateTitle(itemId, title string) error {
	err := i.dbClient.DB().Model(&models.Item{}).Where("item_id = ?", itemId).Update("title", title).Error
	return err
}

func (i *item) UpdateContent(itemId, content string) error {
	err := i.dbClient.DB().Model(&models.Item{}).Where("item_id = ?", itemId).Update("content", content).Error
	return err
}
