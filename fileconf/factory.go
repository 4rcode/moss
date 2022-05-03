package fileconf

import "github.com/4rcode/moss/conf"

// ConfigurerFactory TODO
type ConfigurerFactory Configurer

// NewConfigurer TODO
func (f ConfigurerFactory) NewConfigurer(paths ...string) conf.Configurer {
	f.Paths = append(f.Paths, paths...)

	return Configurer(f)
}
