package app

import (
	"os"
	"strings"
)

func IsDebug() bool {
	if strings.EqualFold(os.Getenv(EnvDebugKey), EnvTrueString) {
		return true
	}

	return false
}
