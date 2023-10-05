package main

import "CompilerPlayground/lexer"

func main() {
	getLexer := lexer.GetLexer("../test/inputUtf8String.txt")
	getLexer.Tokenize()
	getLexer.Print()
}
