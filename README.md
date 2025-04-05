# drasken-go-lexer

A lightweight and extensible lexer written in Go. It tokenizes input text line-by-line, identifying keywords, operators, punctuation, and more — all while supporting custom comment prefixes.

---

## ✨ Features

- Simple and easy to integrate
- Tracks token position (line and column)
- Skips comment lines using customizable prefixes
- Supports:
  - Identifiers and keywords
  - Numeric literals
  - Operators (`=`, `==`, `+`, etc.)
  - Punctuation (`()`, `{}`, `;`, etc.)
- EOF token generation

---

## 📦 Installation

```bash
go get github.com/draskenlabs/drasken-go-lexer@latest
```

## 🔧 Usage

### Import
```go
import "github.com/draskenlabs/drasken-go-lexer/lexer"
```

### Create a lexer and generate tokens
```go
input := `
  // Sample code
  let x = 10
  if x == 10 {
    print(x)
  }
`

commentPrefixes := []string{"//"}
l := lexer.NewLexer(input, commentPrefixes)
tokens := l.GenerateTokens()
```

### Iterate over tokens
```go
for _, tok := range tokens {
	fmt.Printf("Type: %-12s Literal: %-6q Line: %d Col: %d-%d\n",
		tok.Type, tok.Literal, tok.Line, tok.Start, tok.End)
}
```

## 🧪 Output Example

```bash
Type: IDENTIFIER   Literal: "let"   Line: 1 Col: 2-5
Type: IDENTIFIER   Literal: "x"     Line: 1 Col: 6-7
Type: OPERATOR     Literal: "="     Line: 1 Col: 8-9
Type: NUMBER       Literal: "10"    Line: 1 Col: 10-12
Type: IDENTIFIER   Literal: "if"    Line: 2 Col: 2-4
Type: IDENTIFIER   Literal: "x"     Line: 2 Col: 5-6
Type: OPERATOR     Literal: "=="    Line: 2 Col: 7-9
Type: NUMBER       Literal: "10"    Line: 2 Col: 10-12
Type: PUNCTUATION  Literal: "{"     Line: 2 Col: 13-14
...
Type: EOF          Literal: ""      Line: 4 Col: 0-0
```

## 📁 Folder Structure

```bash
.
├── lexer/             # Lexer logic
│   └── lexer.go
├── example/           # Sample usage code
│   └── main.go
├── go.mod
├── README.md
```

## 📄 Token Types

```bash
Token types returned include:

-   `IDENTIFIER`
-   `NUMBER`
-   `OPERATOR`
-   `PUNCTUATION`
-   `EOF`
    
You can expand this for your language or syntax as needed.
```

## 🧰 Example: example/main.go

```go
package main

import (
	"fmt"

	"github.com/draskenlabs/drasken-go-lexer/lexer"
)

func main() {
	input := `
		// This is a comment
		let total = 100
		if total == 100 {
			print(total)
		}
	`

	commentPrefixes := []string{"//"}
	l := lexer.NewLexer(input, commentPrefixes)
	tokens := l.GenerateTokens()

	for _, tok := range tokens {
		fmt.Printf("%-12s %q (Line %d, Col %d-%d)\n",
			tok.Type, tok.Literal, tok.Line, tok.Start, tok.End)
	}
}
```

## 🚀 Roadmap

```bash
-   String literal parsing
-   Unicode support
-   Extended operator recognition
-   Configurable token patterns
```

## 👥 Contributing

We welcome contributions! Please open an issue or submit a PR on [GitHub](https://github.com/draskenlabs/drasken-go-lexer).

## 📄 License

MIT © [Drasken Labs](https://github.com/draskenlabs)


---

Let me know if you also want me to write a `go.mod` or CI workflow (`go test`, etc.) for this project.
