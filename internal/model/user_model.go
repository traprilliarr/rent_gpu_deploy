package model

type NonceResponse struct {
	NonceString string `json:"nonce_id"`
}

type AuthRequest struct {
	PublicAddress string `json:"public_address"`
	Message       string `json:"message"`
	SignedMessage string `json:"signed_message"`
}
