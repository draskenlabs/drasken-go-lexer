package lexer

import (
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
	lines := strings.Split(l.input, "\n")

	for l.linePosition < len(lines) {
		line := lines[l.linePosition]

		// Skip comment lines based on prefixes
		if l.shouldSkipLine(line) {
			l.linePosition++
			continue
		}

		l.columnPosition = 0
		lineLen := len(line)

		for l.columnPosition < lineLen {
			ch := line[l.columnPosition]

			// Skip over whitespace
			if isWhitespace(ch) {
				l.columnPosition++
				continue
			}

			// Handle multi-line backtick strings
			if ch == '`' {
				startLine := l.linePosition
				startColumn := l.columnPosition
				var content strings.Builder
				content.WriteByte('`')
				l.columnPosition++

				terminated := false
				for !terminated && l.linePosition < len(lines) {
					line := lines[l.linePosition]
					for l.columnPosition < len(line) {
						ch := line[l.columnPosition]
						if ch == '`' {
							content.WriteByte('`')
							l.columnPosition++
							terminated = true
							break
						} else {
							content.WriteByte(ch)
							l.columnPosition++
						}
					}
					if !terminated {
						content.WriteString("\n")
						l.linePosition++
						l.columnPosition = 0
					}
				}

				tokens = append(tokens, Token{
					Type:    STRING,
					Literal: content.String(),
					Start:   startColumn,
					End:     l.columnPosition,
					Line:    startLine,
				})

				// Do NOT increment linePosition here — it's handled below the outer loop
				continue
			}

			literalStart := l.columnPosition
			literal := l.generateLiteral(line)
			if strings.TrimSpace(literal) != "" {
				token := GenerateNewToken(literal, literalStart, l.columnPosition, l.linePosition)
				tokens = append(tokens, token)
			}
		}

		l.linePosition++
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

// generateLiteral attempts to extract the next valid token literal from the current line.
// It handles identifiers, numbers, operators (including ==), and punctuation.
func (l *DraskenLexer) generateLiteral(line string) string {
	start := l.columnPosition
	ch := line[start]

	// Handle quoted strings (single or double)
	if ch == '"' || ch == '\'' {
		quote := ch
		i := start + 1
		for i < len(line) {
			if line[i] == quote {
				i++
				break
			}
			i++
		}
		l.columnPosition = i
		return line[start:l.columnPosition]
	}

	// Handle numbers (int or float)
	if isDigit(ch) || (ch == '.' && start+1 < len(line) && isDigit(line[start+1])) {
		i := start
		hasDot := false
		for i < len(line) {
			if line[i] == '.' {
				if hasDot {
					break
				}
				hasDot = true
			} else if !isDigit(line[i]) {
				break
			}
			i++
		}
		l.columnPosition = i
		return line[start:i]
	}

	// Handle identifiers and booleans
	if isLetter(ch) || ch == '_' {
		i := start
		for i < len(line) && isAlphanumericOrUnderscore(line[i]) {
			i++
		}
		l.columnPosition = i
		return line[start:i]
	}

	// Operators
	if isOperator(ch) {
		if ch == '=' && start+1 < len(line) && line[start+1] == '=' {
			l.columnPosition += 2
			return "=="
		}
		l.columnPosition++
		return string(ch)
	}

	// Punctuation
	if isPunctuation(ch) {
		l.columnPosition++
		return string(ch)
	}

	// Fallback for unknowns
	l.columnPosition++
	return string(ch)
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

// shouldSkipLine returns true if the line starts with any comment prefix
func (l *DraskenLexer) shouldSkipLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	for _, prefix := range l.commentPrefixes {
		if strings.HasPrefix(trimmed, prefix) {
			return true
		}
	}
	return false
}
