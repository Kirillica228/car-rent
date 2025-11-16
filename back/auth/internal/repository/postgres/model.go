package postgres



type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	RoleID   uint	`gorm:"not null;default:1"`
	Role     Role   `gorm:"foreignKey:RoleID"`
}

type Role struct {
	ID   uint `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(100);unique;not null"`
}