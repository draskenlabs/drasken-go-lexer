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
```bash
import "github.com/draskenlabs/drasken-go-lexer/lexer"
```