# Playground

Throwaway learning code for the crypto milestones. **Nothing in here is used by the real app.**

- **M2 – crypto primitives** (using `golang.org/x/crypto`):
  1. AEAD: encrypt and decrypt with XChaCha20-Poly1305, then show that tampering breaks decryption.
  2. X25519 key exchange, followed by HKDF key derivation.
  3. Ed25519 signing and verification.
  4. Passphrase key derivation with Argon2id.
- **M3 – Signal protocol**:
  1. A toy symmetric ratchet (educational only).
  2. A `go.mau.fi/libsignal` test in which Alice and Bob build a session and exchange messages.

Suggested layout: one sub-folder per exercise, each with its own `_test.go` so you can run `go test ./playground/...`.

> Rule: home-made crypto is fine *here* for learning. It never goes into the real app.
