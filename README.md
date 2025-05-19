# Cipher API 🔐

A lightweight Go-based microservice that provides secure AES encryption and decryption over HTTP. Ideal for use in IoT, edge processing, or any system requiring centralized, secure encryption logic.

---

## ✨ Features

* AES encryption/decryption (CFB mode)
* Base64 encoding for safe transport
* Secure random IV generation (per request)
* Lightweight and production-ready
* Environment-based key configuration

---

## 📦 API Endpoints

### `POST /encrypt`

Encrypts plain text.

**Request:**

```json
{
  "text": "my secret"
}
```

**Response:**

```json
{
  "result": "encrypted_base64_string"
}
```

### `POST /decrypt`

Decrypts base64-encoded ciphertext.

**Request:**

```json
{
  "text": "encrypted_base64_string"
}
```

**Response:**

```json
{
  "result": "my secret"
}
```

---

## 🔧 Configuration

The service requires an AES key set via an environment variable:

| Variable     | Required | Description                      |
| ------------ | -------- | -------------------------------- |
| `CIPHER_KEY` | ✅ Yes    | Must be 16, 24, or 32 bytes long |

Example:

```bash
export CIPHER_KEY=my16byteSecret!
```

---

## 🐳 Docker Usage

### Build and run locally:

```bash
docker compose up --build -d
```

### Dockerfile summary:

* Multi-stage build (Go → Alpine)
* Secure environment configuration
* Exposes port `8080`

---

## 📁 Folder Structure

```
cipher-api/
├── main.go            # Main application code
├── Dockerfile         # Multi-stage Docker build
├── docker-compose.yaml
└── README.md
```
