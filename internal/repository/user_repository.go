package repository

import (
	"github.com/sirupsen/logrus"
	"rent_gpu_be/internal/entity"
)

type UseriRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUseriRepository(log *logrus.Logger) *UseriRepository {
	return &UseriRepository{
		Log: log,
	}
}
