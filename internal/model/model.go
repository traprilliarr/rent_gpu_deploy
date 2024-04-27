package model

type WebResponse[T any] struct {
	HttpCode int    `json:"http_code"`
	Data     T      `json:"data"`
	Errors   string `json:"errors,omitempty"`
}
