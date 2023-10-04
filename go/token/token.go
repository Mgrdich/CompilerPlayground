package token

import "unicode"

// Token is the set of lexical tokens of our created language
type Token int

// The list of tokens.
const (
	// Special tokens

	ILLEGAL Token = iota
	COMMENT
	EOF

	//  Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	literalBeg
	IDENT
	INTEGER // 12345
	STRING  // "abc"
	literalEnd

	keywordBeg
	VAR // var
	BEGIN
	END
	keywordEnd

	// Operators and delimiters
	operatorBeg
	ADD
	SUB
	MUL
	QUO
	REM
	DEFINE
	DOT
	SEMICOLON
	COMMA
	operatorEnd
)

var tokens = [...]string{

	ILLEGAL: "ILLEGAL",
	COMMENT: "COMMENT",
	EOF:     "EOF",

	IDENT:   "IDENT",
	INTEGER: "INTEGER",
	STRING:  "STRING",

	VAR:   "var",
	BEGIN: "begin",
	END:   "end",

	ADD:       "+",
	SUB:       "-",
	MUL:       "*",
	QUO:       "/",
	REM:       "%",
	DEFINE:    ":=",
	DOT:       ".",
	SEMICOLON: ";",
	COMMA:     ",",
}

var keywords map[string]Token

var predeclaredTypes = map[string]bool{
	"integer": true,
}

func init() {
	keywords = make(map[string]Token, keywordEnd-(keywordBeg+1))
	for i := keywordBeg + 1; i < keywordEnd; i++ {
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
func (tok Token) IsLiteral() bool { return literalBeg < tok && tok < literalEnd }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
func (tok Token) IsOperator() bool {
	return operatorBeg < tok && tok < operatorEnd
}

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
func (tok Token) IsKeyword() bool { return keywordBeg < tok && tok < keywordEnd }

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

// IsPredeclared returns if is predeclared
// lower one work since the map is set , so we don't have to differentiate between
func IsPredeclared(s string) bool {
	return predeclaredTypes[s]
}
