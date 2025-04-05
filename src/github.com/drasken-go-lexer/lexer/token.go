package lexer

import "fmt"

// TokenType defines all possible types of tokens in the lexer.
type TokenType int

const (
	// Special token types that are not valid input values
	ILLEGAL TokenType = iota // Token or character not recognized by the lexer
	EOF                      // End of file/input

	// Literal tokens
	IDENT // Identifier (e.g., variable names, function names)
	INT   // Integer literals (e.g., 42, 1234)

	// Operators
	ASSIGN   // =  (assignment operator)
	PLUS     // +  (addition operator)
	MINUS    // -  (subtraction operator)
	ASTERISK // *  (multiplication operator)
	SLASH    // /  (division operator)
	PERCENT  // %  (modulus operator)
	EQ       // == (equality operator)
	LT       // <  (less than)
	GT       // >  (greater than)
	BANG     // !  (logical negation)
	AND      // &  (bitwise AND)
	OR       // |  (bitwise OR)
	XOR      // ^  (bitwise XOR)

	// Punctuation
	DOT       // .  (dot/period)
	COMMA     // ,  (comma separator)
	SEMICOLON // ;  (semicolon terminator)
	COLON     // :  (colon separator)
	LPAREN    // (  (left parenthesis)
	RPAREN    // )  (right parenthesis)
	LBRACE    // {  (left curly brace)
	RBRACE    // }  (right curly brace)
	LBRACKET  // [  (left square bracket)
	RBRACKET  // ]  (right square bracket)
)

// tokenTypeNames maps TokenType constants to their human-readable string representations.
// This is useful for debugging, logging, and printing tokens.
var tokenTypeNames = [...]string{
	ILLEGAL:   "ILLEGAL",
	EOF:       "EOF",
	IDENT:     "IDENT",
	INT:       "INT",
	ASSIGN:    "ASSIGN",
	PLUS:      "PLUS",
	MINUS:     "MINUS",
	ASTERISK:  "ASTERISK",
	SLASH:     "SLASH",
	PERCENT:   "PERCENT",
	EQ:        "EQ",
	LT:        "LT",
	GT:        "GT",
	BANG:      "BANG",
	AND:       "AND",
	OR:        "OR",
	XOR:       "XOR",
	DOT:       "DOT",
	COMMA:     "COMMA",
	SEMICOLON: "SEMICOLON",
	COLON:     "COLON",
	LPAREN:    "LPAREN",
	RPAREN:    "RPAREN",
	LBRACE:    "LBRACE",
	RBRACE:    "RBRACE",
	LBRACKET:  "LBRACKET",
	RBRACKET:  "RBRACKET",
}

// Token represents a lexical token, which is a single meaningful unit of input.
// Each Token includes metadata about where it was found in the input stream.
type Token struct {
	Type    TokenType // The category/type of the token (e.g., IDENT, PLUS, etc.)
	Literal string    // The actual substring of input that was parsed as this token
	Start   int       // The starting column position of this token in the input line (0-based)
	End     int       // The ending column position of this token in the input line (exclusive)
	Line    int       // The line number in the input where the token was found (0-based)
}

// GenerateNewToken creates a new Token instance by determining its TokenType from the given literal.
// It analyzes the literal string and assigns the appropriate type based on operators, punctuation,
// numeric patterns, or identifier rules. It also sets positional metadata for where the token appears.
func GenerateNewToken(literal string, start int, end int, line int) Token {
	var tokenType TokenType

	switch literal {
	case "+":
		tokenType = PLUS
	case "-":
		tokenType = MINUS
	case "*":
		tokenType = ASTERISK
	case "/":
		tokenType = SLASH
	case "%":
		tokenType = PERCENT
	case "=":
		tokenType = ASSIGN
	case "==":
		tokenType = EQ
	case "<":
		tokenType = LT
	case ">":
		tokenType = GT
	case "!":
		tokenType = BANG
	case "&":
		tokenType = AND
	case "|":
		tokenType = OR
	case "^":
		tokenType = XOR
	case ".":
		tokenType = DOT
	case ",":
		tokenType = COMMA
	case ";":
		tokenType = SEMICOLON
	case ":":
		tokenType = COLON
	case "(":
		tokenType = LPAREN
	case ")":
		tokenType = RPAREN
	case "{":
		tokenType = LBRACE
	case "}":
		tokenType = RBRACE
	case "[":
		tokenType = LBRACKET
	case "]":
		tokenType = RBRACKET
	default:
		// If the literal doesn't match a known operator or punctuation,
		// determine if it's an INT or IDENT based on its first character.
		if isDigit(literal[0]) {
			tokenType = INT
		} else if isLetter(literal[0]) || isUnderscore(literal[0]) {
			tokenType = IDENT
		} else {
			tokenType = ILLEGAL // Anything not recognized is considered illegal.
		}
	}

	return Token{
		Type:    tokenType,
		Literal: literal,
		Start:   start,
		End:     end,
		Line:    line,
	}
}

// String returns a formatted string representation of the Token instance.
// It includes the token type name, its literal value, and the location metadata.
// This is mainly useful for debugging or printing token streams.
func (t Token) String() string {
	return fmt.Sprintf(
		"Type: %s, Literal: '%s', Start: %d, End: %d, Line: %d",
		tokenTypeNames[t.Type], t.Literal, t.Start, t.End, t.Line,
	)
}
