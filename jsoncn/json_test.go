package jsoncn_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/4rcode/moss/boot"
)

func TestConfigurer(t *testing.T) {
	type _data struct {
		Foo    string
		Foobar struct {
			foo, Bar string
		}
	}

	var data _data

	boot.Configurer(json.Configurer{
		json.NewDecoder(strings.NewReader(""))},
	).Configure(&data)

	t.Error(data)
}
