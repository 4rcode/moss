package attest_test

import (
	"testing"

	"github.com/4rcode/moss/attesting/attest"
)

func TestAttest(t *testing.T) {
	is := attest.Error(t).Attest

	is(1 == 2, "no it's not")
}
