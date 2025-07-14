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
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
		{token.COMMA, ","},
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
