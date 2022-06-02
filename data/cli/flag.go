package cli

import (
	"github.com/4rcode/moss/data"
)

type _flag struct {
	build func(string) data.Decoder
	value interface{}
}

func (v _flag) String() string {
	return ""
}

func (v _flag) Set(option string) error {
	return v.
		build(option).
		Decode(v.value)
}
