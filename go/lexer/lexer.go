package lexer

import (
	"CompilerPlayground/token"
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type LexToken struct {
	tok token.Token
	lit string
}

type Lexer struct {
	tokens    []LexToken // TODO think about not storing them
	directory string
	reader    *bufio.Reader
	ch        rune
}

func lower(ch rune) rune     { return ('a' - 'A') | ch } // returns lower-case ch iff ch is ASCII letter
func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }
func isHex(ch rune) bool     { return '0' <= ch && ch <= '9' || 'a' <= lower(ch) && lower(ch) <= 'f' }

func isLetter(ch rune) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return isDecimal(ch) || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
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

	lex.reader = bufio.NewReader(f)
	lex.StartScan()
}

func (lex Lexer) next() (r rune, size int, err error) {
	r, size, err = lex.reader.ReadRune()
	if err == nil {
		lex.ch = r
	}

	return r, size, err
}

func (lex Lexer) StartScan() {
	for {
		tok, lit := lex.Scan()
		if tok == token.EOF {
			break
		}
		lex.tokens = append(lex.tokens, LexToken{tok: tok, lit: lit})
	}
}

func (lex Lexer) Scan() (t token.Token, l string) {
	_, _, err := lex.next()
	if err != nil {
		if err == io.EOF {
			return token.EOF, ""
		}
		panic("something went in the startScan")
	}

	// no error
	var tok token.Token
	var lit string
	switch ch := lex.ch; {
	case isLetter(ch):
		lit = lex.scanIdentifier()
		if len(lit) > 1 {
			// keywords are longer than one letter - avoid lookup otherwise
			tok = token.Lookup(lit)
		} else {
			tok = token.IDENT
		}
	case isDecimal(ch):
		tok, lit = lex.scanNumber()
	default:
		tok = token.ADD
		lit = "testing"
	}

	return tok, lit
}

func (lex Lexer) scanNumber() (t token.Token, lit string) {

}

func (lex Lexer) scanString() string {
	// " should be consumed to start scanning string value
	return ""
}

func (lex Lexer) scanIdentifier() string {
	return ""
}

func (lex Lexer) skipWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' {
		_, _, err := lex.next()
		if err != nil {
			return
		}
	}
}

// Print Mock function that let's us print some values
// to show that the lexer is working normally
func (lex Lexer) Print() {
	for _, lt := range lex.tokens {
		var typing string
		switch {
		case lt.tok == token.ILLEGAL:
			fmt.Println(lt.lit, ": error, invalid lexem")
		case lt.tok.IsNumber():
			typing = "number"
		case lt.tok.IsKeyword() || lt.tok.IsOperator():
			typing = "lexem"
		default:
			typing = "unknown" // TODO something might have gone wrong here
		}

		fmt.Println(typing, ":", lt.lit)
	}
}
