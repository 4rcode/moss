package jsonconf_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/4rcode/moss/boot"
	"github.com/4rcode/moss/jsonconf"
)

func TestConfigurer(t *testing.T) {
	type _data struct {
		Foo    string
		Foobar struct {
			foo, Bar string
		}
	}

	var data _data

	boot.Configurer(jsonconf.Configurer{
		json.NewDecoder(strings.NewReader(""))},
	).Configure(&data)

	t.Error(data)
}
