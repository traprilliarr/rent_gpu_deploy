package model

type OrderRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Telegram string `json:"telegram" validate:"customTelegram"`
	Hash     string `json:"hash"`
	Value    string `json:"value"`
	SSHKEY   string `json:"sshkey" validate:"customSSHKey"`
	GpuID    string `json:"gpu_id"`
	UserID   string `json:"user_id"`
}

type OrderUpdatedRequest struct {
	OrderID string `json:"order_id"`
}

type OrderUserRequest struct {
	UserID string `json:"user_id"`
}

type OrderResponse struct {
	ID                 string `json:"order_id"`
	OrderName          string `json:"gpu_name,omitempty"`
	PaymentAddress     string `json:"payment_address,omitempty"`
	PaymentTransaction string `json:"payment_transaction,omitempty"`
	Value              string `json:"value,omitempty"`
	Status             string `json:"status,omitempty"`
}
