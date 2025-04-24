package modbus

import (
	"time"

	"go.bug.st/serial"
)

// ClientProviderOption client provider option for user.
type ClientProviderOption func(ClientProvider)

// WithLogProvider set logger provider.
func WithLogProvider(provider LogProvider) ClientProviderOption {
	return func(p ClientProvider) {
		p.setLogProvider(provider)
	}
}

// WithEnableLogger enable log output when you has set logger.
func WithEnableLogger() ClientProviderOption {
	return func(p ClientProvider) {
		p.LogMode(true)
	}
}

// WithSerialConfig set serial config, only valid on serial.
func WithSerialConfig(commName string, config serial.Mode) ClientProviderOption {
	return func(p ClientProvider) {
		p.setSerialConfig(commName, config)
	}
}

// WithTimeout set tcp Connect & Read timeout, only valid on TCP.
func WithTimeout(t time.Duration) ClientProviderOption {
	return func(p ClientProvider) {
		p.setTimeout(t)
	}
}
