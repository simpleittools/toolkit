package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-+"

// Tools is the type used to instantiate this module. Any variables of this type will have access
// to all the tools in this module with the receiver *Tools
type Tools struct {
}

// RandomString generates a random string of length n using the characters from randomStringSource.
// it creates a rune then converts it to a string and returns the string
// TODO: provide a usage example
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
