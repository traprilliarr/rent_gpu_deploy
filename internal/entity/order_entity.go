package entity

import "time"

// Order is a struct that represents a user entity
type Order struct {
	ID             string    `gorm:"column:id;primaryKey"`
	Email          string    `gorm:"column:email"`
	Telegram       string    `gorm:"column:telegram"`
	Hash           string    `gorm:"column:hash"`
	Value          string    `gorm:"column:value"`
	PaymentID      string    `gorm:"column:payment_id"`
	PaymentAddress string    `gorm:"column:payment_address"`
	SSHKEY         string    `gorm:"column:ssh_key"`
	GpuID          string    `gorm:"column:gpu_fk"`
	UserID         string    `gorm:"column:user_fk"`
	Status         string    `gorm:"column:status"`
	CreatedAt      time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
}

func (u *Order) TableName() string {
	return "order"
}
