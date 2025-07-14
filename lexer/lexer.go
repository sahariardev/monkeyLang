package lexer

import "monkeyLang/token"

type Lexer struct {
	input        string
	pos          int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

func (l *Lexer) NextToken() token.Token {
	return token.Token{}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	
	l.pos = l.readPosition
	l.readPosition++
}
