package attesting

type Printer[T any] func(T, string, ...any)

// With TODO
func With[L any](logger L, printers ...Printer[L]) {
	for _, print := range printers {
		print(logger, "hi")
	}
}

type Attester func(bool, args ...any) bool

type _attester[L any] struct {
	logger   L
	printers []Printer[L]
}

func (a *_attester[L]) attest() {

}
