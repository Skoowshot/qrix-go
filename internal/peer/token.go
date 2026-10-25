package peer

import (
	"crypto/rand"
	"fmt"
)

// Token is a 128-bit authentication token used to verify a Peer.
type Token [16]byte

// GenerateToken uses crypto random to generate a unique Token.
func GenerateToken() (Token, error) {
	var token Token
	if _, err := rand.Read(token[:]); err != nil {
		return Token{}, fmt.Errorf("failed to generate random token: %w", err)
	}

	return token, nil
}
