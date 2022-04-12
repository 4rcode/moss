package conf

import (
	"bytes"
	"encoding/json"
)

type Decoder interface {
	Decode(interface{}) error
}

type DecoderFactory func([]byte) Decoder

type JsonParser struct {
	Data           []byte
	DecoderFactory DecoderFactory
}

func (p JsonParser) Parse(target interface{}) {
	if p.DecoderFactory == nil {
		p.DecoderFactory = p.bla
	}

	// bla := p.DecoderFactory(p.Data)

	// json.Decoder // json.Unmarshal
}

func (JsonParser) bla(data []byte) Decoder {
	return json.NewDecoder(bytes.NewReader(data))
}
