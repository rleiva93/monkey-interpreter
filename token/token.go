package token

// Type that will be used by the Token structure
type Type string

// Token structure that will be used by the interpreter
type Token struct {
	Type    Type
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
)
