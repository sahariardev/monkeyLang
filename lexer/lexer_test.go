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
	let result = add(five, ten);`

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
