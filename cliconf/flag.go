package cli

import (
	"github.com/4rcode/moss/conf"
)

type _flag struct {
	build func(string) conf.Configurer
	value interface{}
}

func (v _flag) String() string {
	return ""
}

func (v _flag) Set(option string) error {
	return v.
		build(option).
		Configure(v.value)
}
