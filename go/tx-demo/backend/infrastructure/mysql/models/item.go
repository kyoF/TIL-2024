package models

type Item struct {
	ID      int    `gorm:"primaryKey"`
	ItemID  string `gorm:"size:255;unique;not null"`
	Title   string `gorm:"size:255;not null"`
	Content string `gorm:"not null"`
}
