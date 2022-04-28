package conf

import (
	"io"
)

// StringFactory TODO
type StringFactory interface {
	// NewConfigurer TODO
	NewConfigurer(...string) Configurer
}

// ReaderFactory TODO
type ReaderFactory interface {
	// NewConfigurer TODO
	NewConfigurer(io.Reader) Configurer
}
