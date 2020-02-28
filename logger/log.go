package logger

import (
	log "github.com/sirupsen/logrus"
)

// InternalProcessError write internal process errors
func InternalProcessError(err error) {
	log.Error(err)
}
