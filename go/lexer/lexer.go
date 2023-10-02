package lexer

import (
	"bufio"
	"fmt"
	"go/token"
	"os"
)

type tokens struct {
	t     token.Token
	value string
}

type Lexer struct {
	tokens    []tokens
	directory string
	scanner   *bufio.Scanner
}

func GetLexer(directory string) Lexer {
	if len(directory) == 0 {
		panic("directory is not defined")
	}

	return Lexer{directory: directory}
}

func (lex Lexer) Tokenize() {
	if len(lex.directory) == 0 {
		panic("directory is not defined")
	}

	f, err := os.Open(lex.directory)
	if err != nil {
		panic("something wrong with the provided directory")
	}

	lex.scanner = bufio.NewScanner(f)
	lex.startScan()
}

func (lex Lexer) startScan() {
	for lex.scanner.Scan() {
		text := lex.scanner.Text()
		fmt.Println(text)
	}
}

func (lex Lexer) Print() {

}
