package factory

import (
	crp "github.com/kubex-ecosystem/gdbase/internal/security/crypto"
	krs "github.com/kubex-ecosystem/gdbase/internal/security/external"
	sci "github.com/kubex-ecosystem/gdbase/internal/security/interfaces"
)

type CryptoService = sci.ICryptoService
type KeyringService = sci.IKeyringService

func NewCryptoService() CryptoService {
	return crp.NewCryptoService()
}

func NewKeyringService(keyringServiceName, keyringServicePath string) KeyringService {
	return krs.NewKeyringService(keyringServiceName, keyringServicePath)
}
