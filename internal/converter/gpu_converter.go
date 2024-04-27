package converter

import (
	"rent_gpu_be/internal/entity"
	"rent_gpu_be/internal/model"
)

func GpuToResponses(gpus *[]entity.Gpu) []*model.GpuResponse {
	gpuResponses := make([]*model.GpuResponse, 0, len(*gpus))
	for _, gpu := range *gpus {
		gpuResponse := &model.GpuResponse{
			ID:        gpu.ID,
			GpuName:   gpu.GpuName,
			Price:     gpu.Price,
			Available: gpu.Available,
		}
		gpuResponses = append(gpuResponses, gpuResponse)
	}
	return gpuResponses
}

func GpuToResponse(gpus *entity.Gpu) *model.GpuResponse {
	response := model.GpuResponse{
		ID:          gpus.ID,
		GpuName:     gpus.GpuName,
		Price:       gpus.Price,
		Link:        gpus.Link,
		Network:     gpus.Network,
		Cpu:         gpus.Cpu,
		Memory:      gpus.Memory,
		Storage:     gpus.Storage,
		Description: gpus.Description,
		Available:   gpus.Available,
	}
	return &response
}
