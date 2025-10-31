package parser

import (
	"fmt"
	"monkeyLang/ast"
	"monkeyLang/lexer"
	"testing"
)

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserError(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("program.Statements does not contain an return. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral does not match. got=%q", returnStmt.TokenLiteral())
		}
	}

}

func TestLetStatement(t *testing.T) {
	input := `
				let x = 5;
				let y = 6;
				let foobar = 838383;
				`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserError(t, p)

	if program == nil {
		t.Fatalf("ParseProgram returned nil program")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	test := []struct {
		expectedString string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range test {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedString) {
			return
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserError(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	identifier, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("stmt.Expression is not ast.Identifier. got=%T", stmt.Expression)
	}

	if identifier.TokenLiteral() != "foobar" {
		t.Errorf("identifier.TokenLiteral does not match. got=%q", identifier.TokenLiteral())
	}

	if identifier.Value != "foobar" {
		t.Errorf("identifier.Value does not match. got=%q", identifier.Value)
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserError(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("stmt.Expression is not ast.Identifier. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("identifier.TokenLiteral does not match. got=%q", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("identifier.Value does not match. got=%q", literal.TokenLiteral())
	}

}

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserError(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Errorf("exp.Operator does not match. got=%q", exp.Operator)
		}

		if !testIntegerLiteralExpression(t, exp.Right, tt.integerValue) {
			return
		}

	}

}

func testIntegerLiteralExpression(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("il is not ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value does not match. got=%d", integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral does not match. got=%q", integ.TokenLiteral())
		return false
	}

	return true
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral got %s, want let", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s.Statement got %T, want %T", s, ast.LetStatement{})
	}

	if letStatement.Name.Value != name {
		t.Errorf("s.Statement got %s, want %s", letStatement.TokenLiteral(), name)
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("s.Statement got %s, want %s", letStatement.TokenLiteral(), name)
	}

	return true

}

func checkParserError(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, e := range errors {
		t.Errorf("parser error: %s", e)
	}

	t.FailNow()
}
