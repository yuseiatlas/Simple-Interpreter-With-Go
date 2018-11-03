package lexer

import (
	"github.com/yuseiatlas/interpreter/token"
	"testing"
)

func TestNextTokenSimple(t *testing.T) {
	input := `=+()[]{},;.'"`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.DOT, "."},
		{token.SINGLE_QUOTE, "'"},
		{token.DOUBLE_QUOTE, "\""},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, test := range tests {
		myToken := lexer.NextToken()

		if myToken.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, test.expectedType, myToken.Type)
		}

		if myToken.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedLiteral, myToken.Literal)
		}
	}
}

func TestNextTokenComplex(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fun(x, y) {
		x + y;
	};
	let result = add(five, ten);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, test := range tests {
		myToken := lexer.NextToken()

		if myToken.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, test.expectedType, myToken.Type)
		}

		if myToken.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedLiteral, myToken.Literal)
		}
	}
}

func TestNextTokenComplexWithSymbols(t *testing.T) {
	input := `let five = 5;
			let ten = 10.589.5.0..;
			let add = fun(x, y) {
				x + y;
			};
			let result = add(five, ten);
			!-/*5;
			5 < 10 > 5;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10.589"},
		{token.DOT, "."},
		{token.NUMBER, "5.0"},
		{token.DOT, "."},
		{token.DOT, "."},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "5"},
		{token.LT, "<"},
		{token.NUMBER, "10"},
		{token.GT, ">"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, test := range tests {
		myToken := lexer.NextToken()

		if myToken.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, test.expectedType, myToken.Type)
		}

		if myToken.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedLiteral, myToken.Literal)
		}
	}
}

func TestNextTokenComplexWithSymbolsAndBranching(t *testing.T) {
	input := `let five = 5;
			let ten = 10;
			let add = fun(x, y) {
				x + y;
			};
			let result = add(five, ten);
			!-/*5;
			5 < 10 > 5;
			if (5 < 10) {
				return true;
			} else if (5 == 5) {
				return true;
			} else if(5 != 100) {
				return false;
			} else {
				return false;
			}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "5"},
		{token.LT, "<"},
		{token.NUMBER, "10"},
		{token.GT, ">"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.NUMBER, "5"},
		{token.LT, "<"},
		{token.NUMBER, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.NUMBER, "5"},
		{token.EQUAL, "=="},
		{token.NUMBER, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.NUMBER, "5"},
		{token.NOT_EQUAL, "!="},
		{token.NUMBER, "100"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, test := range tests {
		myToken := lexer.NextToken()

		if myToken.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, test.expectedType, myToken.Type)
		}

		if myToken.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedLiteral, myToken.Literal)
		}
	}
}

func TestNextTokenComplexWithLoopsAndStructs(t *testing.T) {
	input := `let five = 5;
			let ten = 10;
			let add = fun(x, y) {
				x + y;
			};
			let result = add(five, ten);
			!-/*5;
			5 < 10 > 5;
			if (5 < 10) {
				return true;
			} else if (5 == 5) {
				return true;
			} else if(5 != 100) {
				return false;
			} else {
				return false;
			}
			while (5 < hello) {
				return false;	
			}
			struct myStruct {
				a: string
				b: double
			}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "5"},
		{token.LT, "<"},
		{token.NUMBER, "10"},
		{token.GT, ">"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.NUMBER, "5"},
		{token.LT, "<"},
		{token.NUMBER, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.NUMBER, "5"},
		{token.EQUAL, "=="},
		{token.NUMBER, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.NUMBER, "5"},
		{token.NOT_EQUAL, "!="},
		{token.NUMBER, "100"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.WHILE, "while"},
		{token.LPAREN, "("},
		{token.NUMBER, "5"},
		{token.LT, "<"},
		{token.IDENT, "hello"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.STRUCT, "struct"},
		{token.IDENT, "myStruct"},
		{token.LBRACE, "{"},
		{token.IDENT, "a"},
		{token.COLON, ":"},
		{token.STRING, "string"},
		{token.IDENT, "b"},
		{token.COLON, ":"},
		{token.DOUBLE, "double"},
		{token.RBRACE, "}"},

		{token.EOF, ""},
	}

	lexer := New(input)

	for i, test := range tests {
		myToken := lexer.NextToken()

		if myToken.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, test.expectedType, myToken.Type)
		}

		if myToken.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedLiteral, myToken.Literal)
		}
	}
}
