// Code generated by berty.tech/core/.scripts/generate-logger.sh

package v0007eventdispatchretry

import "go.uber.org/zap"

func logger() *zap.Logger {
	return zap.L().Named("core.sql.migrations.v0007eventdispatchretry")
}