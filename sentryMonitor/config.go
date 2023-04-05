package sentryMonitor

// SentryMonitoringConfig contains the configuration for the metric collection.
type SentryMonitoringConfig struct {
	Enabled          bool    `toml:",omitempty"`
	DSN              string  `toml:",omitempty"`
	TracesSampleRate float64 `toml:",omitempty"`
	Release          string  `toml:",omitempty"`
}

// DefaultConfig is the default config for metrics used in go-ethereum.
var DefaultConfig = SentryMonitoringConfig{
	Enabled:          false,
	DSN:              "",
	TracesSampleRate: 1.0,
}
