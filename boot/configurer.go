package boot

type Configurer interface {
	Configure(interface{}) error
}
