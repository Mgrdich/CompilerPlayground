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

	INTEGER // 12345
	STRING  // "abc"

	VAR // var
	BEGIN
	END

	ADD
	SUB
	MUL
	QUO
	REM

	DEFINE
)

var tokens = [...]string{

	ILLEGAL: "ILLEGAL",
	COMMENT: "COMMENT",

	INTEGER: "INTEGER",
	STRING:  "STRING",

	VAR:   "VAR",
	BEGIN: "BEGIN",
	END:   "END",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	DEFINE: ":=",
}
