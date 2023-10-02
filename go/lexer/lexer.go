package lexer

import "go/token"

type tokens = struct {
	t     token.Token
	value string
}

type Lexer = struct {
	tokens []tokens
}

func ReadFileCreateLexer(fileDirectory string) {

}

func StdOut(lexerSlice []Lexer) {

}
