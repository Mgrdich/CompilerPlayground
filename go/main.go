package main

import "CompilerPlayground/lexer"

func main() {
	//getLexer := lexer.GetLexer("../test/inputFloat.txt")
	getLexer := lexer.GetLexer("../test/inputUtf8String.txt")
	//getLexer := lexer.GetLexer("../test/input0.txt")
	//getLexer := lexer.GetLexer("../test/input1.txt")
	//getLexer := lexer.GetLexer("../test/input2.txt")
	getLexer.Tokenize()
	getLexer.Print()
}
