goLexerExecFileName:=exec-lexer-go

build-go-lexer:
	go build -o ./$(goLexerExecFileName) ./main.go

run-go-lexer:
	./$(goLexerExecFileName)
