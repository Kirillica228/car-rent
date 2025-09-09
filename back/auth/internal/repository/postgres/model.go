package postgres



type UserModel struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     int    `gorm:"not null;default:0"`
}