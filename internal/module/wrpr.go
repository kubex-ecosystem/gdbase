package module

import (
	"os"
	"strings"
)

func RegX() *GDBase {
	var configPath = os.Getenv("GDBASE_CONFIGFILE")
	var keyPath = os.Getenv("GDBASE_KEYFILE")
	var certPath = os.Getenv("GDBASE_CERTFILE")
	var hideBannerV = os.Getenv("GDBASE_PRINTBANNER")

	return &GDBase{
		configPath: configPath,
		keyPath:    keyPath,
		certPath:   certPath,
		hideBanner: (strings.ToLower(hideBannerV) == "true" ||
			strings.ToLower(hideBannerV) == "1" ||
			strings.ToLower(hideBannerV) == "yes" ||
			strings.ToLower(hideBannerV) == "y"),
	}
}
