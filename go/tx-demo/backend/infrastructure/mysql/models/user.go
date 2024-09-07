package models

type User struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
	Age  int
}
