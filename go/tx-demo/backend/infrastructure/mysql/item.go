package mysqlinfra

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/infrastructure/mysql/models"

	"gorm.io/gorm"
)

type item struct {
	db *gorm.DB
	*transaction
}

func NewItem(db *gorm.DB) repository.Item {
	return &item{
		db:          db,
		transaction: &transaction{db: db},
	}
}

func (i *item) Get(itemId string) (*entity.Item, error) {
	var item models.Item
	err := i.db.Where("item_id = ?", itemId).First(&item).Error
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

	err := i.transaction.txDB.Create(item).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *item) UpdateTitle(itemId, title string) error {
	err := i.transaction.txDB.Model(&models.Item{}).Where("item_id = ?", itemId).Update("title", title).Error
	return err
}

func (i *item) UpdateContent(itemId, content string) error {
	err := i.transaction.txDB.Model(&models.Item{}).Where("item_id = ?", itemId).Update("content", content).Error
	return err
}
