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

// eof is the end of file constant , we can use it during scanIdentifier
const eof = -1

type LexToken struct {
	tok token.Token
	lit string
}

type Lexer struct {
	tokens    []LexToken // TODO think about not storing them
	directory string
	reader    *bufio.Reader
	ch        rune // current rune
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

func GetLexer(directory string) *Lexer {
	if len(directory) == 0 {
		panic("directory is not defined")
	}

	return &Lexer{directory: directory, ch: ' '}
}

// Tokenize will tokenize the whole file , you can actually call lex.Scan() independently as well
func (lex *Lexer) Tokenize() {
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

func (lex *Lexer) next() (r rune, size int) {
	r, size, err := lex.reader.ReadRune()
	if err == nil {
		lex.ch = r
	}

	if err == io.EOF {
		lex.ch = eof
	}

	return r, size
}

func (lex *Lexer) peekRune() rune {
	for peekBytes := 4; peekBytes > 0; peekBytes-- { // unicode rune can be up to 4 bytes
		b, err := lex.reader.Peek(peekBytes)
		if err == nil {
			decodeRune, _ := utf8.DecodeRune(b)
			if decodeRune == utf8.RuneError {
				return utf8.RuneError
			}

			// success
			return decodeRune
		}
		// Otherwise, we ignore Peek errors and try the next smallest number of bytes
	}

	// Pretty sure we can assume EOF if we get this far
	return 0
}

func (lex *Lexer) StartScan() {
	for {
		tok, lit := lex.Scan()
		if tok == token.EOF {
			break
		}
		lex.tokens = append(lex.tokens, LexToken{tok: tok, lit: lit})
	}
}

func (lex *Lexer) Scan() (t token.Token, l string) {
	// Someone from before reads already reach the end of file
	if lex.ch == eof {
		return token.EOF, ""
	}

	lex.skipWhitespace()

	// no error
	var tok token.Token
	var lit string
	switch ch := lex.ch; {
	case ch == eof:
		tok = token.EOF
		lit = ""
	case isLetter(ch):
		lit = lex.scanIdentifier()
		if len(lit) > 1 {
			// keywords are longer than one letter - avoid lookup otherwise
			tok = token.Lookup(lit)
		} else {
			tok = token.IDENT
		}
	case isDecimal(ch) || ch == '.' && isDecimal(lex.peekRune()):
		tok, lit = lex.scanNumber()
	default:
		lex.next()
		switch {
		case ch == '"':
			// cause this might return invalid string
			tok, lit = lex.scanString()
		case ch == ',':
			tok = token.COMMA
			lit = ","
		case ch == '+':
			tok = token.ADD
			lit = "+"
		case ch == '-':
			tok = token.SUB
			lit = "-"
		case ch == '*':
			tok = token.MUL
			lit = "*"
		case ch == '/':
			tok = token.QUO
			lit = "/"
		case ch == '%':
			tok = token.REM
			lit = "%"
		case ch == ':':
			tok = lex.switch2(token.COLON, token.DEFINE)
			if tok == token.DEFINE {
				lit = ":="
			} else {
				lit = ":"
			}
		case ch == ';':
			tok = token.SEMICOLON
			lit = ";"
		case ch == '.':
			tok = token.DOT
			lit = "."
		default:
			tok = token.ILLEGAL
			lit = string(ch)
		}
	}

	return tok, lit
}

// digits can only read base 10 for now
// even though the go source code implements it as an arbitrary base
func (lex *Lexer) digits(builtNumber *[]rune, state token.Token) token.Token {
	tokenState := state
	m := rune('0' + 10)
	for isDecimal(lex.ch) {
		if lex.ch >= m {
			tokenState = token.ILLEGAL
		}
		*builtNumber = append(*builtNumber, lex.ch)
		lex.next()
	}
	return tokenState
}

func (lex *Lexer) scanNumber() (tok token.Token, lit string) {
	tok = token.ILLEGAL
	builtNumber := []rune{lex.ch} // TODO research where we can keep it as byte

	lex.next()

	if lex.ch != '.' {
		tok = token.INTEGER
	}

	if lex.ch == '.' {
		builtNumber = append(builtNumber, lex.ch) // add the dot notation
		lex.next()
		tok = token.FLOAT
	}

	tok = lex.digits(&builtNumber, tok)

	return tok, string(builtNumber)
}

func (lex *Lexer) scanString() (tok token.Token, lit string) {
	var builtWord []rune // because this is consumed and nexted already

	for {
		ch := lex.ch
		if ch == '\n' || ch < 0 {
			return token.ILLEGAL, string(builtWord)
		}

		lex.next()
		if ch == '"' {
			break
		} else {
			builtWord = append(builtWord, ch)
		}
	}

	return token.STRING, string(builtWord)
}

func (lex *Lexer) scanIdentifier() string {
	// add already nexted letter cause in order this process to start is should be a letter
	builtWord := []rune{lex.ch}

	lex.next()
	for isLetter(lex.ch) || isDigit(lex.ch) {
		// this position is important cause it will make eof case work
		builtWord = append(builtWord, lex.ch)
		lex.next()
	}

	return string(builtWord)
}

func (lex *Lexer) skipWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.next()
	}
}

// Helper functions for scanning multi-byte tokens such as >> += >>= .
// Different routines recognize different length tok_i based on matches
// of ch_i. If a token ends in '=', the result is tok1 or tok3
// respectively. Otherwise, the result is tok0 if there was no other
// matching character, or tok2 if the matching character was ch2.

func (lex *Lexer) switch2(tok0, tok1 token.Token) token.Token {
	if lex.ch == '=' {
		lex.next()
		return tok1
	}
	return tok0
}

// Print Mock function that let's us print some values
// to show that the lexer is working normally
func (lex *Lexer) Print() {
	for _, lt := range lex.tokens {
		var typing string
		switch {
		case lt.tok == token.ILLEGAL:
			fmt.Println(lt.lit + ": error, invalid lexem")
		case lt.tok.IsNumber():
			typing = "number"
		case lt.tok == token.STRING:
			typing = "string"
		case lt.tok == token.IDENT:
			typing = "ident"
		case lt.tok.IsKeyword() || lt.tok.IsOperator():
			typing = "lexem"
		default:
			typing = "unknown" // TODO something might have gone wrong here
		}

		if lt.tok != token.ILLEGAL {
			fmt.Println(typing + ":\"" + lt.lit + "\"")
		}
	}
}
