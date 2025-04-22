package handlers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/rednexx46/cipher-api/models"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func init() {
	if len(secretKey) != 32 {
		panic("SECRET_KEY must be 32 bytes long")
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.IndexResponse{Message: "Welcome to the Cipher API!"})
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.HealthCheckResponse{Status: "OK"})
}

func EncryptHandler(w http.ResponseWriter, r *http.Request) {
	var req models.EncryptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	cipherText, err := encrypt(req.PlainText)
	if err != nil {
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.EncryptResponse{CipherText: cipherText})
}

func DecryptHandler(w http.ResponseWriter, r *http.Request) {
	var req models.DecryptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	plainText, err := decrypt(req.CipherText)
	if err != nil {
		http.Error(w, "Decryption failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.DecryptResponse{PlainText: plainText})
}

func encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func decrypt(encryptedText string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", err
	}

	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
