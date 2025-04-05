package lexer

import (
	"bufio"
	"fmt"
	"strings"
)

// DraskenLexer is the core lexer structure that holds the input string,
// current character position, and configuration such as comment prefixes.
type DraskenLexer struct {
	input           string   // the full input source code to be tokenized
	position        int      // byte offset in the entire input string (not used here but can be useful for future extensions)
	ch              byte     // current character under examination (not used here but reserved for extensions)
	commentPrefixes []string // list of valid comment prefixes used to skip comment lines
	columnPosition  int      // column index within the current line
	linePosition    int      // line index in the input
}

// NewLexer creates and returns a new instance of DraskenLexer given an input string
// and comment prefixes. It initializes the lexer without scanning yet.
// This sets up the lexer structure with the provided source input and comment prefix list,
// but does not yet begin the process of tokenization.
func NewLexer(input string, commentPrefixes []string) *DraskenLexer {
	l := &DraskenLexer{input: input, commentPrefixes: commentPrefixes}
	return l
}

// GenerateTokens reads the input line-by-line and generates a slice of Tokens.
// It uses bufio.Scanner for line-based iteration and delegates token generation
// to generateLineTokens. At the end, it appends an EOF token.
func (l *DraskenLexer) GenerateTokens() []Token {
	tokens := []Token{}

	scanner := bufio.NewScanner(strings.NewReader(l.input))
	for scanner.Scan() {
		l.columnPosition = 0
		tok := l.generateLineTokens(scanner.Text())
		if tok != nil {
			tokens = append(tokens, tok...)
		}
		l.linePosition++
	}

	// Handle potential scanner error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
	}

	// Append EOF token at the end of token stream
	tokens = append(tokens, Token{
		Type:    EOF,
		Literal: "",
		Start:   l.columnPosition,
		End:     l.columnPosition,
		Line:    l.linePosition,
	})
	return tokens
}

// generateLineTokens parses a single line of input into tokens by repeatedly
// extracting literals using generateLiteral, skipping over whitespace.
func (l *DraskenLexer) generateLineTokens(line string) []Token {
	tokens := []Token{}
	lineLen := len(line)

	for l.columnPosition < lineLen {
		ch := line[l.columnPosition]

		// Skip over whitespace
		if isWhitespace(ch) {
			l.columnPosition++
			continue
		}

		literalStart := l.columnPosition
		literal := l.generateLiteral(line)
		if strings.TrimSpace(literal) != "" {
			token := GenerateNewToken(literal, literalStart, l.columnPosition, l.linePosition)
			tokens = append(tokens, token)
		}
	}

	return tokens
}

// generateLiteral attempts to extract the next valid token literal from the current line.
// It handles identifiers, numbers, operators (including ==), and punctuation.
// Comment prefix detection is also performed to skip irrelevant lines.
func (l *DraskenLexer) generateLiteral(line string) string {
	var literal string
	i := l.columnPosition

	for i < len(line) {
		ch := line[i]

		// Check if the current character starts a known comment prefix
		isKnownComment := true
		for _, prefix := range l.commentPrefixes {
			if strings.HasPrefix(strings.TrimSpace(string(ch)), prefix) {
				isKnownComment = false
				break
			}
		}
		if !isKnownComment {
			i = len(line) // skip rest of the line
			break
		} else if isWhitespace(ch) {
			break
		} else if isAlphanumericOrUnderscore(ch) {
			literal += string(ch) // build identifier or number
		} else if literal != "" && (isOperator(ch) || isPunctuation(ch)) {
			break // stop if operator or punctuation follows a literal
		} else if isOperator(ch) {
			// Detect double-character operators (e.g., ==)
			if ch == '=' && i+1 < len(line) && line[i+1] == '=' {
				literal += "=="
				i += 2
			} else {
				literal += string(ch)
				i++
			}
			break
		} else if isPunctuation(ch) {
			literal += string(ch)
			i++
			break
		}
		i++
	}

	l.columnPosition = i
	return literal
}

// Helper function to check if a character is a letter (A-Z, a-z)
func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// Helper function to check if a character is a digit (0-9)
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// isOperator returns true if the character is one of the supported operator symbols.
func isOperator(ch byte) bool {
	switch ch {
	case '+', '-', '*', '/', '%', '=', '<', '>', '!', '&', '|', '^':
		return true
	default:
		return false
	}
}

// isPunctuation returns true if the character is one of the punctuation symbols.
func isPunctuation(ch byte) bool {
	switch ch {
	case '.', ',', ';', ':', '(', ')', '{', '}', '[', ']':
		return true
	default:
		return false
	}
}

// isUnderscore returns true if the character is an underscore.
func isUnderscore(ch byte) bool {
	return ch == '_'
}

// isWhitespace returns true if the character is a whitespace (space, tab, newline, carriage return).
func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// isAlphanumericOrUnderscore checks whether a character is a letter, digit, or underscore.
func isAlphanumericOrUnderscore(ch byte) bool {
	return isLetter(ch) || isDigit(ch) || isUnderscore(ch)
}
