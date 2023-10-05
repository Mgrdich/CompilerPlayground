package main

import "CompilerPlayground/lexer"

func main() {
	getLexer := lexer.GetLexer("../test/input0.txt")
	getLexer.Tokenize()
	getLexer.Print()
}
