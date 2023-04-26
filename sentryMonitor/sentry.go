package sentryMonitor

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
)

var config SentryMonitoringConfig

func StartMonitoring(cfg SentryMonitoringConfig) error {
	config = cfg

	err := sentry.Init(sentry.ClientOptions{
		Dsn: cfg.DSN,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: cfg.TracesSampleRate,
		EnableTracing:    true,
		Release:          cfg.Release,
	})
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.LevelError)
	})

	return err
}

func HandleError(err error) {
	if !config.Enabled {
		return
	}

	sentry.CaptureException(err)
}

func HandleErrorWithArgs(msg string, args ...interface{}) {
	if !config.Enabled {
		return
	}
	err := msg + fmt.Sprint(args...)
	sentry.CaptureException(errors.New(err))
}

func HandleWarningWithArgs(msg string, args ...interface{}) {
	if !config.Enabled {
		return
	}
	warn := msg + fmt.Sprint(args...)
	sentry.CaptureMessage(warn)
}

func HandleFatal(format string, args ...interface{}) {
	if !config.Enabled {
		return
	}
	err := fmt.Sprintf("Fatal: "+format+"\n", args...)
	sentry.CaptureException(errors.New(err))
}

func HandleErrorMessage(err string) {
	if !config.Enabled {
		return
	}
	sentry.CaptureException(errors.New(err))
}

func HandleMessage(message string) {
	if !config.Enabled {
		return
	}
	sentry.CaptureMessage(message)
}
