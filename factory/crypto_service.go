package factory

import (
	crp "github.com/rafa-mori/gdbase/internal/security/crypto"
	krs "github.com/rafa-mori/gdbase/internal/security/external"
	sci "github.com/rafa-mori/gdbase/internal/security/interfaces"
)

type CryptoService = sci.ICryptoService
type KeyringService = sci.IKeyringService

func NewCryptoService() CryptoService {
	return crp.NewCryptoService()
}

func NewKeyringService(keyringServiceName, keyringServicePath string) KeyringService {
	return krs.NewKeyringService(keyringServiceName, keyringServicePath)
}
