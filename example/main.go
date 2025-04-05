package main

import (
	"fmt"
	"os"

	"github.com/drasken/drasken-go-lexer/src/github.com/drasken-go-lexer/lexer"
)

func main() {
	dat, err := os.ReadFile("./files/first.xyz")
	if err != nil {
		panic(err)
	}

	lexer := lexer.NewLexer(string(dat), []string{"#"})
	tokens := lexer.GenerateTokens()
	fmt.Println(tokens)
}
