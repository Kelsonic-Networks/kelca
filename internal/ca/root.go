package ca

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"github.com/kelsonic-networks/kelca/internal/crypto"
	"github.com/kelsonic-networks/kelca/internal/storage"
)

type RootCA struct {
	CommonName   string
	Organization string
	KeyType      string
	KeySize      int
	Validity     int
}

func (ca *RootCA) Create(store *storage.SecureStorage) error {
	priv, err := crypto.GenerateKey(ca.KeyType, ca.KeySize)
	if err != nil {
		return err
	}

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return err
	}

	notBefore := time.Now()
	notAfter := notBefore.AddDate(0, 0, ca.Validity)

	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   ca.CommonName,
			Organization: []string{ca.Organization},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            0,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, template, template, priv.Public(), priv)
	if err != nil {
		return err
	}

	return store.SaveCertificate("root", derBytes)
}
