package main

import "CompilerPlayground/lexer"

func main() {
	//getLexer := lexer.GetLexer("./tests/inputFloat.txt")
	getLexer := lexer.GetLexer("./tests/inputUtf8String.txt")
	//getLexer := lexer.GetLexer("./tests/input0.txt")
	//getLexer := lexer.GetLexer("./tests/input1.txt")
	//getLexer := lexer.GetLexer("./tests/input2.txt")
	getLexer.Tokenize()
	getLexer.Print()
}
