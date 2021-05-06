package TokenType

type TokenTypes int

const (
	Plus  TokenTypes = iota // +
	Minus                   // -
	Star                    // *
	Slash                   // /

	GE // >=
	GT // >
	EQ // ==
	LE // <=
	LT // <

	SemiColon  // ;
	LeftParen  // (
	RightParen // )

	Assignment // =

	If
	Else

	Int

	Identifier //标识符

	IntLiteral //整型字面量
	StringLiteral
)

func (d TokenTypes) String() string {
	return [...]string{"Plus", "Minus", "Star", "Slash", "GE", "GT", "EQ", "LE", "LT", "SemiColon", "LeftParen", "RightParen", "Assignment", "If", "Else", "Int", "Identifier", "IntLiteral", "StringLiteral"}[d]
}
