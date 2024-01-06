package service

import (
	"fmt"

	siwe "github.com/spruceid/siwe-go"
)

var (
	siweErr     error
	siweMessage *siwe.Message
)

func AuthGenerateNonceUseCase() string {
	nonce := siwe.GenerateNonce()

	return nonce
}

func AuthVerifyMessageUseCase(signature string, message string, nonce *string) error {
	siweMessage, siweErr = siwe.ParseMessage(message)

	if siweErr != nil {
		return siweErr
	}

	_, errValid := siweMessage.ValidNow()

	if errValid != nil {
		return errValid
	}

	_, siweErr = siweMessage.Verify(signature, nil, nonce, nil)

	if siweErr != nil {
		return siweErr
	}

	address := siweMessage.GetAddress()

	fmt.Println(address)

	return nil
}
