package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Signifies a token/character we don't know about

	EOF = "EOF" // Tells the parser when to stop

	// Identifiers + literals
	IDENT  = "IDENT"  // add, foobar, x, y, ...
	NUMBER = "NUMBER" // 1343456

	// Operators
	ASSIGN    = "="
	EQUAL     = "=="
	NOT_EQUAL = "!="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"

	LT = "<"
	GT = ">"

	// Separators
	COMMA        = ","
	DOT          = "."
	COLON        = ":"
	SEMICOLON    = ";"
	SINGLE_QUOTE = "'"
	DOUBLE_QUOTE = "\""

	// Delimeters
	LPAREN   = "("
	RPAREN   = ")"
	LBRACKET = "["
	RBRACKET = "]"
	LBRACE   = "{"
	RBRACE   = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	WHILE    = "WHILE"
	STRUCT   = "STRUCT"
	STRING   = "string"
	DOUBLE   = "DOUBLE"
	INT      = "INT"
)

var keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"while":  WHILE,
	"struct": STRUCT,
	"string": STRING,
	"double": DOUBLE,
	"int":    INT,
}

func GetIdentifierType(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
