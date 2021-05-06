package tokenReader

type Token struct {
	tokenType string
	tokenText string
	offfset   int
}

type SimpleToken struct {
	Token
	Type string
	Text string
}

type TokenReader struct {
	tokens []SimpleToken
	pos    int
}

func (r *TokenReader) Read() SimpleToken {
	if r.pos < len(r.tokens) {
		var result SimpleToken = r.tokens[r.pos]
		r.pos = r.pos + 1
		return result
	}
	return SimpleToken{}
}

func (r *TokenReader) Peek() SimpleToken {
	if r.pos < len(r.tokens) {
		return r.tokens[r.pos]
	}
	return SimpleToken{}
}

func (r *TokenReader) Unread() {
	if r.pos > 0 {
		r.pos -= 1
	}
}

func (r *TokenReader) GetPosition() int {
	return r.pos
}

func (r *TokenReader) SetPosition(position int) {
	if position >= 0 && position < len(r.tokens) {
		r.pos = position
	}
}
