package controller

import (
	"encoding/hex"
	"net/http"

	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
}

func (_self MessageHandler) SomeoneCreate(c *gin.Context) {
	var messRequest SomeOneCreateMessRequest
	err := c.Bind(&messRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Decode the hex-encoded public key of the recipient.
	pubKeyBytes, err := hex.DecodeString(messRequest.PublicKey) // uncompressed pub key
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	pubKey, err := secp256k1.ParsePubKey(pubKeyBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Encrypt a message decrypted by the private key corresponding to pubKey
	ciphertext, err := secp256k1.Encrypt(pubKey, []byte(messRequest.Text))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"succeed":           true,
		"encrypted_message": hex.EncodeToString(ciphertext),
	})
}

func (_self MessageHandler) ICreate(c *gin.Context) {
	var messRequest ICreateMessRequest
	err := c.Bind(&messRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Decode a hex-encoded private key.
	pkBytes, err := hex.DecodeString(messRequest.PrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	privateKey, _ := secp256k1.PrivKeyFromBytes(pkBytes)

	// Sign a message using the private key
	messageBytes := []byte(messRequest.Text)
	signature, err := privateKey.Sign(messageBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"succeed":      true,
		"signature":    hex.EncodeToString(signature.Serialize()),
		"message_hash": hex.EncodeToString(messageBytes),
	})

}

func (_self MessageHandler) GenKeys(c *gin.Context) {
	randomPrivateKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"succeed":               true,
		"private_key":           hex.EncodeToString(randomPrivateKey.Serialize()),
		"public_key_compressed": hex.EncodeToString(randomPrivateKey.PubKey().SerializeCompressed()),
	})
}

func (_self MessageHandler) DecryptMessage(c *gin.Context) {
	var decryptMessageRequest DecryptMessRequest
	err := c.Bind(&decryptMessageRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Decode the hex-encoded private key.
	pkBytes, err := hex.DecodeString(decryptMessageRequest.PrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	privateKey, _ := secp256k1.PrivKeyFromBytes(pkBytes)
	ciphertext, err := hex.DecodeString(decryptMessageRequest.EncryptedMessage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Try decrypting the message.
	plaintext, err := secp256k1.Decrypt(privateKey, ciphertext)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"succeed": true,
		"message": string(plaintext),
	})
}

func (_self MessageHandler) VerifyMessage(c *gin.Context) {
	var verifyMessRequest VerifyMessRequest
	err := c.Bind(&verifyMessRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	signatureBytes, err := hex.DecodeString(verifyMessRequest.Signature)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	messageBytes, err := hex.DecodeString(verifyMessRequest.EncryptedMessage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	publicKeyBytes, err := hex.DecodeString(verifyMessRequest.PublicKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	publicKey, err := secp256k1.ParsePubKey(publicKeyBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	signature, err := secp256k1.ParseSignature(signatureBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !signature.Verify(messageBytes, publicKey) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid signature",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"succeed":  true,
		"verified": true,
		"message":  string(messageBytes),
	})
}
