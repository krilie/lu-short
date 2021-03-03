package rsautil

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

func ParseRsaPublicKey(pubPEMData string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEMData))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

func ParseRsaPrivateKey(pemstr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemstr))
	if block == nil || len(block.Bytes) == 0 {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) // pkcs v 15
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func ParseRsaPublicKeyFromBase64(bytes string) (*rsa.PublicKey, error) {
	block, err := base64.StdEncoding.DecodeString(bytes)
	if err != nil {
		return nil, err
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block) // pkix ?
	if err != nil {
		return nil, errors.New("parse public key error:" + err.Error())
	}
	pubkey, _ := pubInterface.(*rsa.PublicKey)
	return pubkey, nil
}

func ParseRsaPrivateKeyFromBase64(bytes string) (*rsa.PrivateKey, error) {
	block, err := base64.StdEncoding.DecodeString(bytes)
	if err != nil {
		return nil, err
	}
	priv, err := x509.ParsePKCS8PrivateKey(block) // pkcs v 15
	if err != nil {
		return nil, err
	}
	prikey, _ := priv.(*rsa.PrivateKey)
	return prikey, nil
}
