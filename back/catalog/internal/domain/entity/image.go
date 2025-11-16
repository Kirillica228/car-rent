package entity

type Image struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	CarID uint   `gorm:"index"`
	URL   string `gorm:"not null"`
}