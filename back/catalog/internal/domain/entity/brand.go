package entity

type Brand struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"unique;not null" json:"name"`
	IsVisible bool   `gorm:"default:false" json:"is_visible"`
}