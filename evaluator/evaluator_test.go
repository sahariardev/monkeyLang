package evaluator

import (
	"monkeyLang/lexer"
	"monkeyLang/object"
	"monkeyLang/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"100", 100},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBoolean(t, evaluated, tt.expected)
	}

}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)

	if !ok {
		t.Errorf("object is not Integer. obj: %v", obj)
	}

	if result.Value != expected {
		t.Errorf("result.Value: %v, expected: %v", result.Value, expected)
		return false
	}

	return true
}

func testBoolean(t *testing.T, evaluated object.Object, expected bool) {
	result, ok := evaluated.(*object.Boolean)

	if !ok {
		t.Errorf("object is not Boolean. obj: %v", evaluated)
	}

	if result.Value != expected {
		t.Errorf("result.Value: %v, expected: %v", result.Value, expected)
	}
}
