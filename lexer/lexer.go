package lexer

import (
	"errors"
	"github.com/selectiveduplicate/cosmoc/token"
	"strconv"
)

// Interpreter is the lexer struct
type Interpreter struct {
	text         string
	position     int
	currentToken *token.Token
	currentChar  string
}

// NewInterpreter returns a new Interpreter instance
func NewInterpreter(text string, position int, currentToken *token.Token, currentChar string) Interpreter {
	return Interpreter{text, position, currentToken, currentChar}
}

// IsNumeric determines whether the current char
// is a number or not
func (i *Interpreter) IsNumeric() bool {
	if _, err := strconv.Atoi(i.currentChar); err != nil {
		return false
	}
	return true
}

////////// Methods for Interpreter ///////////////

// RaiseError method should need be
func (i *Interpreter) RaiseError() error {
	return errors.New("Error parsing input")
}

// Advance advances position along the input
func (i *Interpreter) Advance() {
	i.position++
	if i.position > len(i.text)-1 {
		i.currentChar = ""
	} else {
		i.currentChar = string(i.text[i.position])
	}
}

// MakeStrInt returns a stringified multi-digit integer
func (i *Interpreter) MakeStrInt() string {
	var strResult string
	for i.currentChar != "" && i.IsNumeric() {
		strResult += i.currentChar
		i.Advance()
	}
	return strResult
}

// GetNextToken spits out the next Token
func (i *Interpreter) GetNextToken() *token.Token {
	for i.currentChar != "" {
		if i.currentChar == "+" {
			i.Advance()
			token := token.Token{Type: token.PLUS, Value: i.currentChar}
			return &token
		}
		if i.IsNumeric() {
			token := token.Token{Type: token.INTEGER, Value: i.MakeStrInt()}
			return &token
		}
	}
	i.RaiseError()
	return nil
}

// Eat compares the currentToken with the passed one
// and calls to get the next token
func (i *Interpreter) Eat(t *token.Token) {
	if i.currentToken.Type == t.Type {
		i.currentToken = i.GetNextToken()
	} else {
		i.RaiseError()
	}
}

// Expression evaluates the expression
func (i *Interpreter) Expression() int {
	i.currentToken = i.GetNextToken()

	left := i.currentToken
	i.Eat(left)

	operator := i.currentToken
	i.Eat(operator)

	right := i.currentToken
	i.Eat(right)

	leftValue, _ := strconv.Atoi(left.Value)
	rightValue, _ := strconv.Atoi(right.Value)

	return leftValue + rightValue
}