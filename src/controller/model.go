package controller

type AuthVerifyMessageModel struct {
	Domain    string `json:"domain" validate:"required"`
	Message   string `json:"message" validate:"required"`
	Signature string `json:"signature" validate:"required"`
}
