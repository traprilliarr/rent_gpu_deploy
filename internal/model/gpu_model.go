package model

type GpuResponse struct {
	ID          string `json:"id"`
	GpuName     string `json:"gpu_name"`
	Price       string `json:"price"`
	Link        string `json:"link,omitempty"`
	Network     string `json:"network,omitempty"`
	Cpu         string `json:"cpu,omitempty"`
	Memory      string `json:"memory,omitempty"`
	Storage     string `json:"storage,omitempty"`
	Description string `json:"description,omitempty"`
	Available   bool   `json:"available,omitempty"`
}
