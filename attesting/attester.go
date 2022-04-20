package attesting

// Attester TODO
type Attester []func(...interface{})

// Attest TODO
func (a Attester) Attest(truth bool, args ...interface{}) bool {
	if truth {
		return true
	}

	for _, log := range a {
		if log != nil {
			log(args...)
		}
	}

	return false
}
