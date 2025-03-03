package client

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"golang.org/x/crypto/pkcs12"
)

// loadPKCS12 loads a PKCS12 encoded certificate.
// Parameters:
//   - data: the PKCS12 encoded certificate bytes
//   - password: the password for the PKCS12 encoded certificate
//
// Returns:
//   - tls.Certificate: The loaded certificate
//   - error: An error if loading fails
func loadPKCS12(data []byte, password string) (tls.Certificate, error) {
	// Try the more flexible PEM conversion first
	blocks, err := pkcs12.ToPEM(data, password)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed converting pkcs12 to PEM: %w", err)
	}

	var tlsCert tls.Certificate
	var privateKey crypto.PrivateKey

	// Extract certificates and private key from PEM blocks
	for _, block := range blocks {
		switch block.Type {
		case "CERTIFICATE":
			tlsCert.Certificate = append(tlsCert.Certificate, block.Bytes)
		case "PRIVATE KEY":
			if privateKey != nil {
				return tls.Certificate{}, fmt.Errorf("found multiple private keys in PKCS12 file")
			}
			var err error
			privateKey, err = parsePrivateKey(block.Bytes)
			if err != nil {
				return tls.Certificate{}, fmt.Errorf("failed parsing private key: %w", err)
			}
		}
	}

	if len(tlsCert.Certificate) == 0 {
		return tls.Certificate{}, fmt.Errorf("no certificates found in PKCS12 file")
	}
	if privateKey == nil {
		return tls.Certificate{}, fmt.Errorf("no private key found in PKCS12 file")
	}

	tlsCert.PrivateKey = privateKey
	return tlsCert, nil
}

// parsePrivateKey attempts to parse a private key in PKCS#1 or SEC1 format
func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	// Try various private key formats
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, nil
	}
	return nil, fmt.Errorf("failed to parse private key in common formats")
}

// tLSConfigForPKCS12 creates a tls.Config from a PKCS12 encoded certificate.
// Parameters:
//   - data: the PKCS12 encoded certificate bytes
//   - password: the password for the PKCS12 encoded certificate
//
// Returns:
//   - *tls.Config: TLS configuration using the certificate
//   - error: An error if configuration fails
func tLSConfigForPKCS12(data []byte, password string) (*tls.Config, error) {
	cert, err := loadPKCS12(data, password)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		Certificates:  []tls.Certificate{cert},
		Renegotiation: tls.RenegotiateOnceAsClient,
	}, nil
}
