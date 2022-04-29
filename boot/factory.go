package boot

import (
	"io"
)

type StringFactory[C Configurer] interface {
	NewConfigurer(...string) C
}

type ReaderFactory[C Configurer] interface {
	NewConfigurer(io.Reader) C
}
