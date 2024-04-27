package repository

import (
	"github.com/sirupsen/logrus"
	"rent_gpu_be/internal/entity"
)

type GpuRepository struct {
	Repository[entity.Gpu]
	Log *logrus.Logger
}

func NewGpuRepository(log *logrus.Logger) *GpuRepository {
	return &GpuRepository{
		Log: log,
	}
}
