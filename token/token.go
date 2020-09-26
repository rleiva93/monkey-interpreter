package token

// Type that will be used by the Token structure
type TokenType string

// Token structure that will be used by the interpreter
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL expressions that are not recognized
	ILLEGAL = "ILLEGAL"
	// EOF indicates the end of hte file
	EOF = "EOF"

	// IDENT will represent the identation spaces
	IDENT = "IDENT"
	// INT corresponds to the type of variable
	INT = "INT"

	// ASSIGN is the operator to bind one value to its identifier
	ASSIGN = "="
	// PLUS sign for addition and concatenation
	PLUS = "+"
	// MINUS sign for substraction
	MINUS = "-"
	// BANG sign for negation
	BANG = "!"
	// ASTERISK for multiplication
	ASTERISK = "*"
	// SLASH for division
	SLASH = "/"
	// LT less than
	LT = "<"
	// GT greater than
	GT = ">"
	// EQ equals
	EQ = "=="
	// NOTEQ different
	NOTEQ = "!="

	// COMMA delimites two elements within the same expression
	COMMA = ","
	// SEMICOLON represents the end of the file
	SEMICOLON = ";"

	// LPAREN opening parenthesis char
	LPAREN = "("
	// RPAREN closing parenthesis char
	RPAREN = ")"
	// LBRACE opening curly brace char
	LBRACE = "{"
	// RBRACE closing curly brace char
	RBRACE = "}"

	// FUNCTION keyword
	FUNCTION = "FUNCTION"
	// LET keyword
	LET = "LET"
	// TRUE keyword
	TRUE = "TRUE"
	// FALSE keyword
	FALSE = "FALSE"
	// IF keyword
	IF = "IF"
	// ELSE keyword
	ELSE = "ELSE"
	// RETURN keyword
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks for indentation chars
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
