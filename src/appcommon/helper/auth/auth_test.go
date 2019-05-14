package auth

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	GenerateToken("acc")
	// TODO: Add parse token function to parse the generated token
}
