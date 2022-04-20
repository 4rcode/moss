package file

import "github.com/4rcode/moss/conf"

// ConfigurerFactory TODO
type ConfigurerFactory Configurer

func (c ConfigurerFactory) NewConfigurer(paths ...string) conf.Configurer {
	c.Paths = append(c.Paths, paths...)

	return (*Configurer)(&c)
}
