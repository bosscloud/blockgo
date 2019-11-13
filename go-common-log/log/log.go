package log

import (
	"github.com/huangbosbos/blockgo/go-common-log/logging"
)

var (
	logger = logging.New(nil, nil, new(logging.ColouredFormatter))

	// INFO ...
	INFO = logger[logging.INFO]
	// WARNING ...
	WARNING = logger[logging.WARNING]
	// ERROR ...
	ERROR = logger[logging.ERROR]
	// FATAL ...
	FATAL = logger[logging.FATAL]
)

// Set sets a custom logger
func Set(l logging.LoggerInterface) {
	INFO = l
	WARNING = l
	ERROR = l
	FATAL = l
}
