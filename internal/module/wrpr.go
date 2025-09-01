package module

import (
	"os"
	"strings"
)

func RegX() *GDBase {
	var configPath = os.Getenv("GODOBASE_CONFIGFILE")
	var keyPath = os.Getenv("GODOBASE_KEYFILE")
	var certPath = os.Getenv("GODOBASE_CERTFILE")
	var printBannerV = os.Getenv("GODOBASE_PRINTBANNER")
	if printBannerV == "" {
		printBannerV = "true"
	}

	return &GDBase{
		configPath:  configPath,
		keyPath:     keyPath,
		certPath:    certPath,
		printBanner: strings.ToLower(printBannerV) == "true",
	}
}
