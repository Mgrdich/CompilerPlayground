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
	numBeg
	INTEGER // 12345
	FLOAT
	numEnd // 123.323
	STRING // "abc"
	literalEnd

	keywordBeg
	VAR // var
	BEGIN
	END
	IntKeyword
	FloatKeyword
	StringKeyword
	keywordEnd

	// Operators and delimiters
	operatorBeg
	ADD
	SUB
	MUL
	QUO
	REM
	DEFINE
	COLON
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
	FLOAT:   "FLOAT",
	STRING:  "STRING",

	VAR:           "var",
	BEGIN:         "begin",
	END:           "end",
	IntKeyword:    "integer",
	FloatKeyword:  "float",
	StringKeyword: "string",

	ADD:       "+",
	SUB:       "-",
	MUL:       "*",
	QUO:       "/",
	REM:       "%",
	DEFINE:    ":=",
	DOT:       ".",
	SEMICOLON: ";",
	COLON:     ":",
	COMMA:     ",",
}

var keywords map[string]Token

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

// IsNumber returns true for tokens corresponding to a numbers
func (tok Token) IsNumber() bool {
	return numBeg < tok && tok < numEnd
}

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
func (tok Token) IsKeyword() bool { return keywordBeg < tok && tok < keywordEnd }

// IsKeyword reports whether name is a our language keyword, such as "func" or "return".
func IsKeyword(name string) bool {
	_, ok := keywords[name]
	return ok
}

// IsIdentifier reports whether name is a our language identifier, that is, a non-empty
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
