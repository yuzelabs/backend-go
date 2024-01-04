package service

import (
	siwe "github.com/spruceid/siwe-go"
)

func AuthGenerateNonceUseCase() string {
	nonce := siwe.GenerateNonce()

	return nonce
}

func AuthVerifyMessageUseCase() {

}
