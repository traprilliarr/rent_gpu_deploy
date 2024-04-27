package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"rent_gpu_be/internal/entity"
)

type OrderRepository struct {
	Repository[entity.Order]
	Log *logrus.Logger
}

func NewOrderRepository(log *logrus.Logger) *OrderRepository {
	return &OrderRepository{
		Log: log,
	}
}
func (r *Repository[T]) FindAllByUserId(db *gorm.DB, entity *[]T, id any) error {
	return db.Where("user_fk = ?", id).Find(entity).Error
}
