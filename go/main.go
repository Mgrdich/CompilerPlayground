package main

import "CompilerPlayground/lexer"

func main() {
	getLexer := lexer.GetLexer("../test/inputUtf8.txt")
	getLexer.Tokenize()
	getLexer.Print()
}
