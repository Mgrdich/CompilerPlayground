package token

import "unicode"

// Token is the set of lexical tokens of our created language
type Token int

// The list of tokens.
const (
	// Special tokens

	ILLEGAL Token = iota
	COMMENT

	//  Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	literal_beg
	IDENT
	INTEGER // 12345
	STRING  // "abc"
	literal_end

	keyword_beg
	VAR // var
	BEGIN
	END
	keyword_end

	// Operators and delimiters
	operator_beg
	ADD
	SUB
	MUL
	QUO
	REM
	DEFINE
	DOT
	SEMICOLOUMN
	operator_end
)

var tokens = [...]string{

	ILLEGAL: "ILLEGAL",
	COMMENT: "COMMENT",

	IDENT:   "IDENT",
	INTEGER: "INTEGER",
	STRING:  "STRING",

	VAR:   "VAR",
	BEGIN: "BEGIN",
	END:   "END",

	ADD:         "+",
	SUB:         "-",
	MUL:         "*",
	QUO:         "/",
	REM:         "%",
	DEFINE:      ":=",
	DOT:         ".",
	SEMICOLOUMN: ";",
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token, keyword_end-(keyword_beg+1))
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
func Lookup(ident string) Token {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	return IDENT
}

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
func (tok Token) IsLiteral() bool { return literal_beg < tok && tok < literal_end }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
func (tok Token) IsOperator() bool {
	return operator_beg < tok && tok < operator_end
}

// IsKeyword reports whether name is a Go keyword, such as "func" or "return".
func IsKeyword(name string) bool {
	_, ok := keywords[name]
	return ok
}

// IsIdentifier reports whether name is a Go identifier, that is, a non-empty
// string made up of letters, digits, and underscores, where the first character
// is not a digit. Keywords are not identifiers.
func IsIdentifier(name string) bool {
	if name == "" || IsKeyword(name) {
		return false
	}
	for i, c := range name {
		if !unicode.IsLetter(c) && c != '_' && (i == 0 || !unicode.IsDigit(c)) {
			return false
		}
	}
	return true
}
