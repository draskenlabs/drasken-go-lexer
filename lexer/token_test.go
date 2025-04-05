package lexer

import (
	"testing"
)

// TestGenerateNewToken verifies that GenerateNewToken correctly identifies
// token types based on various input literals, and that it sets the correct
// position metadata (start, end, and line).
func TestGenerateNewToken(t *testing.T) {
	// Define a list of test cases to run
	tests := []struct {
		literal string    // the string input to be tokenized
		start   int       // the start position of the token
		end     int       // the end position (exclusive)
		line    int       // the line number in the source
		want    TokenType // expected token type
	}{
		// Operators
		{"+", 0, 1, 0, PLUS},
		{"-", 1, 2, 0, MINUS},
		{"*", 2, 3, 0, ASTERISK},
		{"/", 3, 4, 0, SLASH},
		{"%", 4, 5, 0, PERCENT},
		{"=", 5, 6, 0, ASSIGN},
		{"==", 6, 8, 0, EQ},
		{"<", 8, 9, 0, LT},
		{">", 9, 10, 0, GT},
		{"!", 10, 11, 0, BANG},
		{"&", 11, 12, 0, AND},
		{"|", 12, 13, 0, OR},
		{"^", 13, 14, 0, XOR},

		// Punctuation
		{".", 14, 15, 0, DOT},
		{",", 15, 16, 0, COMMA},
		{";", 16, 17, 0, SEMICOLON},
		{":", 17, 18, 0, COLON},
		{"(", 18, 19, 0, LPAREN},
		{")", 19, 20, 0, RPAREN},
		{"{", 20, 21, 0, LBRACE},
		{"}", 21, 22, 0, RBRACE},
		{"[", 22, 23, 0, LBRACKET},
		{"]", 23, 24, 0, RBRACKET},

		// Literals
		{"123", 0, 3, 1, INT},
		{"abc", 4, 7, 1, IDENT},
		{"_foo", 8, 12, 1, IDENT},

		// Illegal token
		{"@", 13, 14, 1, ILLEGAL},
	}

	// Run each test case
	for _, tt := range tests {
		tok := GenerateNewToken(tt.literal, tt.start, tt.end, tt.line)

		// Check token type
		if tok.Type != tt.want {
			t.Errorf("GenerateNewToken(%q) returned Type %v, want %v", tt.literal, tok.Type, tt.want)
		}

		// Check token metadata
		if tok.Literal != tt.literal || tok.Start != tt.start || tok.End != tt.end || tok.Line != tt.line {
			t.Errorf("GenerateNewToken(%q) returned wrong metadata: %+v", tt.literal, tok)
		}
	}
}

// TestTokenString checks the output of the Token.String() method to ensure
// it returns the expected formatted string containing type, literal, and position info.
func TestTokenString(t *testing.T) {
	// Create a sample token
	tok := Token{
		Type:    IDENT,
		Literal: "foo",
		Start:   1,
		End:     4,
		Line:    2,
	}

	// Expected string output
	want := "Type: IDENT, Literal: 'foo', Start: 1, End: 4, Line: 2"

	// Validate output of String() method
	if got := tok.String(); got != want {
		t.Errorf("Token.String() = %q, want %q", got, want)
	}
}
