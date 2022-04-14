package attest

import (
	"testing"

	"github.com/4rcode/moss/attesting"
)

func Error(tb testing.TB) attesting.Attester {
	return attesting.Attester{tb.Error}
}

func Fatal(tb testing.TB) attesting.Attester {
	return attesting.Attester{tb.Fatal}
}

func Log(tb testing.TB) attesting.Attester {
	return attesting.Attester{tb.Log}
}

func Skip(tb testing.TB) attesting.Attester {
	return attesting.Attester{tb.Skip}
}
