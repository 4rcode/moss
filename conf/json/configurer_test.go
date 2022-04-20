package json_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/4rcode/moss/conf"
	json_conf "github.com/4rcode/moss/conf/json"
)

func TestConfigurer(t *testing.T) {
	type _data struct {
		Foo    string
		bar    int
		Foobar struct {
			foo, Bar string
		}
	}

	var data _data

	conf.Configurer(json_conf.Configurer{
		json.NewDecoder(strings.NewReader(""))},
	).Configure(&data)

	t.Error(data)
}
