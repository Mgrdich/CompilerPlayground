package main

import "CompilerPlayground/lexer"

func main() {
	getLexer := lexer.GetLexer("../test/inputFloat.txt")
	getLexer.Tokenize()
	getLexer.Print()
}
