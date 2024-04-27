package entity

import "time"

// Gpu is a struct that represents a user entity
type Gpu struct {
	ID          string    `gorm:"column:id;primaryKey"`
	GpuName     string    `gorm:"column:gpu_name"`
	Price       string    `gorm:"column:price"`
	Link        string    `gorm:"column:link"`
	Network     string    `gorm:"column:network"`
	Cpu         string    `gorm:"column:cpu"`
	Memory      string    `gorm:"column:memory"`
	Storage     string    `gorm:"column:storage"`
	Description string    `gorm:"column:description"`
	Available   bool      `gorm:"column:available"`
	Order       []Order   `gorm:"foreignKey:gpu_fk"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
}

func (u *Gpu) TableName() string {
	return "gpu"
}
