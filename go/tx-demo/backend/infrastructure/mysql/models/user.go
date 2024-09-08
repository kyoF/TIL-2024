package models

type User struct {
	ID     int    `gorm:"primaryKey"`
	UserID string `gorm:"size:255;unique;not null"`
	Name   string `gorm:"size:255;not null"`
	Age    int    `gorm:"not null"`
}
