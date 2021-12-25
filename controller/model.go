package controller

type SomeOneCreateMessRequest struct {
	PublicKey string `json:"public_key"`
	Text      string `json:"text"`
}

type ICreateMessRequest struct {
	PrivateKey string `json:"private_key"`
	Text       string `json:"text"`
}

type DecryptMessRequest struct {
	PrivateKey       string `json:"private_key"`
	EncryptedMessage string `json:"encrypted_message"`
}

type VerifyMessRequest struct {
	EncryptedMessage string `json:"encrypted_message"`
	Signature        string `json:"signature"`
	PublicKey        string `json:"public_key"`
}
