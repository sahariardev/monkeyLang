package evaluator

import (
	"monkeyLang/ast"
	"monkeyLang/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatement(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	}

	return nil
}

func evalStatement(stmt []ast.Statement) object.Object {
	var result object.Object
	for _, stmt := range stmt {
		result = Eval(stmt)
	}

	return result
}
