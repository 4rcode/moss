package prop

type Decoder interface {
	Decode(interface{}) error
}
