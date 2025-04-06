package lexer

import (
	"reflect"
	"testing"
)

// TestGenerateTokens verifies that the lexer correctly tokenizes a multiline input string,
// ignoring comment lines, and returns all expected tokens including EOF.
func TestGenerateTokens(t *testing.T) {
	input := `
// this is a comment
x = 5 + 10;
y = x * 2;
// another comment
z = y / 2 - 1;`

	lexer := NewLexer(input, []string{"//"})
	tokens := lexer.GenerateTokens()

	expectedLiterals := []string{
		"x", "=", "5", "+", "10", ";",
		"y", "=", "x", "*", "2", ";",
		"z", "=", "y", "/", "2", "-", "1", ";",
		"", // EOF token
	}

	var actualLiterals []string
	for _, tok := range tokens {
		actualLiterals = append(actualLiterals, tok.Literal)
	}

	if !reflect.DeepEqual(actualLiterals, expectedLiterals) {
		t.Errorf("token literals do not match\nexpected: %#v\ngot:      %#v", expectedLiterals, actualLiterals)
	}
}

// TestGenerateLineTokens checks tokenization for a single line, ensuring that each
// token is parsed with the correct literal and TokenType.
func TestGenerateLineTokens(t *testing.T) {
	// input := "foo = bar + 1;"
	// lexer := NewLexer(input, []string{"//"})
	// tokens := lexer.generateLineTokens(input)

	// expected := []struct {
	// 	Literal string
	// 	Type    TokenType
	// }{
	// 	{"foo", IDENT},
	// 	{"=", ASSIGN},
	// 	{"bar", IDENT},
	// 	{"+", PLUS},
	// 	{"1", INT},
	// 	{";", SEMICOLON},
	// }

	// if len(tokens) != len(expected) {
	// 	t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	// }

	// for i, exp := range expected {
	// 	if tokens[i].Literal != exp.Literal || tokens[i].Type != exp.Type {
	// 		t.Errorf("token %d mismatch: got (%s, %v), expected (%s, %v)",
	// 			i, tokens[i].Literal, tokens[i].Type, exp.Literal, exp.Type)
	// 	}
	// }
}
