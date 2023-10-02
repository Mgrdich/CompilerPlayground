package token

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
	INTEGER // 12345
	STRING  // "abc"
	literal_end

	keyword_beg
	VAR // var
	BEGIN
	END
	DOT
	keyword_end

	operator_beg
	ADD
	SUB
	MUL
	QUO
	REM
	DEFINE
	operator_end
)

var tokens = [...]string{

	ILLEGAL: "ILLEGAL",
	COMMENT: "COMMENT",

	INTEGER: "INTEGER",
	STRING:  "STRING",

	VAR:   "VAR",
	BEGIN: "BEGIN",
	END:   "END",
	DOT:   ".",

	ADD:    "+",
	SUB:    "-",
	MUL:    "*",
	QUO:    "/",
	REM:    "%",
	DEFINE: ":=",
}
