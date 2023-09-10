package lexer

import "interpreter/token"

type Lexer struct {
	input        string
	position     int  // 所输入字符串中的当前位置 （指向当前字符）
	readPosition int  // 所输入字符串中的当前读取位置 指向当前字符之后的一个字符
	ch           byte // 当前正在查看的字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar() 读取input中的下一个字符，并前移其在input中的位置
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // 检查是否到达字符串末尾
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NewToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	}
}
