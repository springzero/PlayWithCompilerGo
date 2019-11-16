package main

import "fmt"

type ss int

const (
	Plus  ss = iota // +
	Minus           // -
	Star            // *
	Slash           // /

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

func (d ss) String() string {
	return [...]string{"Plus", "Minus", "Star", "Slash", "GE", "GT", "EQ", "LE", "LT", "SemiColon", "LeftParen", "RightParen", "Assignment", "If", "Else", "Int", "Identifier", "IntLiteral", "StringLiteral"}[d]
}

func main() {
	// fmt.Print("start")
	// var tmp byte = '4'
	// //var ss string = ""
	// fmt.Println(tmp > '0')
	// fmt.Println(tmp < '9')
	fmt.Print("start")

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	fmt.Printf("end1 \n")

	fmt.Printf("end2 \n")
	fmt.Printf("end3 \n")
}
