// main.go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Payload struct {
	Text string `json:"text"`
}

type Response struct {
	Result string `json:"result"`
}

var secretKey []byte

func encrypt(text string) (string, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(enc string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	var p Payload
	json.NewDecoder(r.Body).Decode(&p)
	result, err := encrypt(p.Text)
	if err != nil {
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Response{Result: result})
}

func decryptHandler(w http.ResponseWriter, r *http.Request) {
	var p Payload
	json.NewDecoder(r.Body).Decode(&p)
	result, err := decrypt(p.Text)
	if err != nil {
		http.Error(w, "Decryption failed", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Response{Result: result})
}

func main() {
	key := os.Getenv("CIPHER_KEY")
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		log.Fatal("CIPHER_KEY must be 16, 24 or 32 bytes")
	}
	secretKey = []byte(key)

	http.HandleFunc("/encrypt", encryptHandler)
	http.HandleFunc("/decrypt", decryptHandler)

	fmt.Println("[Cipher API] Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
