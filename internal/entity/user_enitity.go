package entity

import "time"

// User is a struct that represents a user entity
type User struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Hash      string    `gorm:"column:hash"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
}

func (u *User) TableName() string {
	return "user"
}
