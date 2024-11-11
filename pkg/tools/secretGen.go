package tools

import (
	"crypto/md5"
	"encoding/hex"
)

// SecretKey variável para armazenar a chave criptografada
var SecretKey = ""

// CreateHash função para criptografar o texto limpo
func CreateHash(key string) (err error) {
	// iniciando o modulo de md5
	hasher := md5.New()
	// transformando a string para byte e escrevendo o hash
	hasher.Write([]byte(key))
	// retornando o hash em md5
	SecretKey = hex.EncodeToString(hasher.Sum(nil))

	return
}

// GetSecretKey função para retornar a chave criptografada
func GetSecretKey() (secretKey string) {
	return SecretKey
}
