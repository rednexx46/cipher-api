## 🔐 Cipher API – Encrypt & Decrypt Service

A simple HTTP API to encrypt and decrypt text using AES-GCM.

---

### 🌐 Base URL

```
http://localhost:8080
```

---

## 📌 Endpoints

### 🔒 POST `/encrypt`

Encrypts plain text using AES-GCM.

#### ✅ Request Body

```json
{
  "plain_text": "Hello World"
}
```

#### 🔁 Response

```json
{
  "cipher_text": "BASE64_ENCRYPTED_STRING"
}
```

#### 🔧 Notes

- The `cipher_text` is base64 encoded and includes a nonce prefix for AES-GCM decryption.

---

### 🔓 POST `/decrypt`

Decrypts a base64-encoded ciphertext string.

#### ✅ Request Body

```json
{
  "cipher_text": "BASE64_ENCRYPTED_STRING"
}
```

#### 🔁 Response

```json
{
  "plain_text": "Hello World"
}
```

#### ⚠️ Errors

- `400 Bad Request` → if input is invalid
- `500 Internal Server Error` → if decryption fails or ciphertext is malformed

---

## 🔐 Environment Variables

| Variable     | Description              | Required | Example            |
| ------------ | ------------------------ | -------- | ------------------ |
| `SECRET_KEY` | AES key (16/24/32 bytes) | ✅       | `examplekey123456` |

---

## 🧪 Quick Test with `curl`

**Encrypt:**

```bash
curl -X POST http://localhost:8080/encrypt \
  -H "Content-Type: application/json" \
  -d '{"plain_text": "test message"}'
```

**Decrypt:**

```bash
curl -X POST http://localhost:8080/decrypt \
  -H "Content-Type: application/json" \
  -d '{"cipher_text": "<base64_string>"}'
```

---

Let me know if you want to turn this into Swagger/OpenAPI docs or generate Postman collections too.
