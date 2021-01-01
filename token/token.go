package token

import "fmt"

const (
	INTEGER = "INTEGER"
	PLUS    = "PLUS"
	EOF     = "EOF"
)

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

// Methods for Token
// string representation of the token
func (t *Token) Str() string {
	return fmt.Sprintf("Token(%v, %v)", t.Type, t.Value)
}
