package attest

import (
	"testing"

	"github.com/4rcode/moss/attesting"
)

func Error(tb testing.TB) attesting.Attester {
	return attesting.Attester{tb.Errorf}
}

func Fatal(tb testing.TB) attesting.Attester {
	return attesting.Attester{tb.Fatalf}
}

func Log(tb testing.TB) attesting.Attester {
	return attesting.Attester{tb.Logf}
}

func Skip(tb testing.TB) attesting.Attester {
	return attesting.Attester{tb.Skipf}
}
