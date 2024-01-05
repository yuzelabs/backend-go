package controller

type AuthVerifyMessageModel struct {
	Message   string `json:"message" validate:"required"`
	Signature string `json:"signature" validate:"required"`
}
