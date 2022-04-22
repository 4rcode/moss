package conf

import (
	"io"
)

// StringConfigurerFactory TODO
type StringConfigurerFactory interface {
	// NewConfigurer TODO
	NewConfigurer(...string) Configurer
}

// ReaderConfigurerFactory TODO
type ReaderConfigurerFactory interface {
	// NewConfigurer TODO
	NewConfigurer(io.Reader) Configurer
}
