package lexer

import (
	"bytes"
	"fmt"

	"github.com/springzero/PlayWithCompilerGo/03/TokenType"
)

//---token define
type Token struct {
	tokenType string
	tokenText string
	offfset   int
}

func (t *Token) getType() string {
	return t.tokenType
}

func (t *Token) getText() string {
	return t.tokenText
}

type SimpleToken struct {
	Token
	Type string
	Text string
}

//----

// func main() {
// 	var lex SimpleLexer = SimpleLexer{}

// 	var script string = "int age = 45;"
// 	var tokenReader SimpleTokenReader = lex.Tokenize(script)
// 	dump(tokenReader, script)

// 	script = "inta age = 45;"
// 	tokenReader = lex.Tokenize(script)
// 	dump(tokenReader, script)

// 	script = "in age = 45;"
// 	tokenReader = lex.Tokenize(script)
// 	dump(tokenReader, script)

// 	script = "age >= 45;"
// 	tokenReader = lex.Tokenize(script)
// 	dump(tokenReader, script)

// }

type SimpleLexer struct {
	tokenText string
	tokens    []SimpleToken
	token     SimpleToken
}

func (lex *SimpleLexer) isAlpha(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func (lex *SimpleLexer) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (lex *SimpleLexer) isBlank(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

/* 有限状态机进入初始状态 */
func (lex *SimpleLexer) initToken(ch byte) DfaState {
	if len(lex.tokenText) > 0 {
		lex.token.Text = lex.tokenText
		lex.tokens = append(lex.tokens, lex.token)

		lex.tokenText = ""
		lex.token = SimpleToken{}
	}
	var newState DfaState = Initial
	if lex.isAlpha(ch) { //第一个字符是字母
		if ch == 'i' {
			newState = Id_int1
		} else {
			newState = Id //进入Id状态
		}
		lex.token.Type = TokenType.Identifier.String()
		lex.tokenText += string(ch)
	} else if lex.isDigit(ch) { //第一个字符是数字
		newState = IntLiteral
		lex.token.Type = TokenType.IntLiteral.String()
		lex.tokenText += string(ch)
	} else if ch == '>' { //第一个字符是>
		newState = GT
		lex.token.Type = TokenType.GT.String()
		lex.tokenText += string(ch)
	} else if ch == '+' {
		newState = Plus
		lex.token.Type = TokenType.Plus.String()
		lex.tokenText += string(ch)
	} else if ch == '-' {
		newState = Minus
		lex.token.Type = TokenType.Minus.String()
		lex.tokenText += string(ch)
	} else if ch == '*' {
		newState = Star
		lex.token.Type = TokenType.Star.String()
		lex.tokenText += string(ch)
	} else if ch == '/' {
		newState = Slash
		lex.token.Type = TokenType.Slash.String()
		lex.tokenText += string(ch)
	} else if ch == ';' {
		newState = SemiColon
		lex.token.Type = TokenType.SemiColon.String()
		lex.tokenText += string(ch)
	} else if ch == '(' {
		newState = LeftParen
		lex.token.Type = TokenType.LeftParen.String()
		lex.tokenText += string(ch)
	} else if ch == ')' {
		newState = RightParen
		lex.token.Type = TokenType.RightParen.String()
		lex.tokenText += string(ch)
	} else if ch == '=' {
		newState = Assignment
		lex.token.Type = TokenType.Assignment.String()
		lex.tokenText += string(ch)
	} else {
		newState = Initial
	}
	return newState
}

func (lex SimpleLexer) Tokenize(code string) SimpleTokenReader {
	lex.tokens = []SimpleToken{}
	reader := bytes.NewBufferString(code)
	lex.tokenText = ""
	lex.token = SimpleToken{}
	var ich byte = 0
	var ch byte = 0
	var state DfaState = Initial
	for {
		ich, _ = reader.ReadByte()
		if ich == 0 {
			break
		}
		ch = ich
		switch state {
		case Initial:
			state = lex.initToken(ch)
		case Id:
			if lex.isAlpha(ch) || lex.isDigit(ch) {
				lex.tokenText += string(ch)
			} else {
				state = lex.initToken(ch)
			}
		case GT:
			if ch == '=' {
				lex.token.Type = TokenType.GE.String()
				state = GE
				lex.tokenText += string(ch)
			} else {
				state = lex.initToken(ch)
			}
		case GE:
			state = lex.initToken(ch)
		case Assignment:
			state = lex.initToken(ch)
		case Plus:
			state = lex.initToken(ch)
		case Minus:
			state = lex.initToken(ch)
		case Star:
			state = lex.initToken(ch)
		case Slash:
			state = lex.initToken(ch)
		case SemiColon:
			state = lex.initToken(ch)
		case LeftParen:
			state = lex.initToken(ch)
		case RightParen:
			state = lex.initToken(ch)
		case IntLiteral:
			if lex.isDigit(ch) {
				lex.tokenText += string(ch)
			} else {
				state = lex.initToken(ch)
			}
		case Id_int1:
			if ch == 'n' {
				state = Id_int2
				lex.tokenText += string(ch)
			} else if lex.isDigit(ch) || lex.isAlpha(ch) {
				state = Id
				lex.tokenText += string(ch)
			} else {
				state = lex.initToken(ch)
			}
		case Id_int2:
			if ch == 't' {
				state = Id_int3
				lex.tokenText += string(ch)
			} else if lex.isDigit(ch) || lex.isAlpha(ch) {
				state = Id
				lex.tokenText += string(ch)
			} else {
				state = lex.initToken(ch)
			}
		case Id_int3:
			if lex.isBlank(ch) {
				lex.token.Type = TokenType.Int.String()
				state = lex.initToken(ch)
			} else {
				state = Id
				lex.tokenText += string(ch)
			}
		default:

		}

	}
	if len(lex.tokenText) > 0 {
		lex.initToken(ch)
	}
	return SimpleTokenReader{tokens: lex.tokens}
}

type SimpleTokenReader struct {
	tokens []SimpleToken
	pos    int
}

func (r *SimpleTokenReader) Read() SimpleToken {
	if r.pos < len(r.tokens) {
		var result SimpleToken = r.tokens[r.pos]
		r.pos = r.pos + 1
		return result
	}
	return SimpleToken{}
}

func (r *SimpleTokenReader) Peek() SimpleToken {
	if r.pos < len(r.tokens) {
		return r.tokens[r.pos]
	}
	return SimpleToken{}
}

func (r *SimpleTokenReader) Unread() {
	if r.pos > 0 {
		r.pos -= 1
	}
}

func (r *SimpleTokenReader) GetPosition() int {
	return r.pos
}

func (r *SimpleTokenReader) SetPosition(position int) {
	if position >= 0 && position < len(r.tokens) {
		r.pos = position
	}
}

func dump(tokenReader SimpleTokenReader, script string) {

	var token SimpleToken
	fmt.Printf("script: %s\n", script)
	for {
		token = tokenReader.Read()
		if token.Text == "" {
			break
		}
		fmt.Printf("type: %s \t text: %s\n", token.Type, token.Text)
	}
	fmt.Println()

}

type DfaState int

const (
	Initial DfaState = iota

	If
	Id_if1
	Id_if2
	Else
	Id_else1
	Id_else2
	Id_else3
	Id_else4
	Int
	Id_int1
	Id_int2
	Id_int3
	Id
	GT
	GE

	Assignment

	Plus
	Minus
	Star
	Slash

	SemiColon
	LeftParen
	RightParen

	IntLiteral
)

func (d DfaState) String() string {
	return [...]string{"Initial", "If", "Id_if1", "Id_if2", "Else", "Id_else1", "Id_else2", "Id_else3", "Id_else4",
		"Int", "Id_int1", "Id_int2", "Id_int3", "Id", "GT", "GE",
		"Assignment", "Plus", "Minus", "Star", "Slash",
		"SemiColon", "LeftParen", "RightParen", "IntLiteral"}[d]
}
