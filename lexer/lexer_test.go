package lexer

import (
	"monkeyLang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		currentToken := l.NextToken()

		if currentToken.Type != tt.expectedType {
			t.Fatalf("test[%d], expected=%q, got=%q", i,
				tt.expectedType, currentToken.Type)
		}

		if currentToken.Literal != tt.expectedValue {
			t.Fatalf("test[%d], expected=%s, got=%s", i,
				tt.expectedValue, currentToken.Literal)
		}
	}
}

func TestNextTokenDetail(t *testing.T) {
	input := `let five = 5;
	 let ten = 10;
	 let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	!-*/<>
	if(5 < 10) {
		return true;
    } else {
		return false;
	}
	
	10 == 10;
    10 != 9;
`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
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
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQUAL, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOTEQUAL, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		currentToken := l.NextToken()

		if currentToken.Type != tt.expectedType {
			t.Fatalf("test[%d], expected=%q, got=%q", i,
				tt.expectedType, currentToken.Type)
		}

		if currentToken.Literal != tt.expectedValue {
			t.Fatalf("test[%d], expected=%s, got=%s", i,
				tt.expectedValue, currentToken.Literal)
		}
	}

}
