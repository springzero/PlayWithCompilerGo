package main

import (
	"fmt"
	"senzflow/PlayWithCompilerGo/three/ASTNode"
	"senzflow/PlayWithCompilerGo/three/ASTNodeType"
	"senzflow/PlayWithCompilerGo/three/TokenType"
	"senzflow/PlayWithCompilerGo/three/lexer"
	"strconv"
)

type SimpleCalculator struct {
}

func main() {
	var calculator SimpleCalculator = SimpleCalculator{}

	var script string = "int a = b+3;"
	fmt.Println("解析变量声明语句： " + script)
	var lex lexer.SimpleLexer = lexer.SimpleLexer{}
	var tokens lexer.SimpleTokenReader = lex.Tokenize(script)

	var node ASTNode.ASTNode = calculator.intDeclare(&tokens)
	calculator.dumpAST(node, "")

	script = "2+3*5"
	fmt.Println("计算：" + script + ", 看上去一切正常。")
	calculator.evaluateOne(script)
}

func (c *SimpleCalculator) evaluateOne(script string) {
	var tree ASTNode.ASTNode = c.Parse(script)
	c.dumpAST(tree, "")
	c.evaluate(tree, "")
}

func (c *SimpleCalculator) Parse(code string) ASTNode.ASTNode {
	var lex lexer.SimpleLexer = lexer.SimpleLexer{}
	var tokens lexer.SimpleTokenReader = lex.Tokenize(code)

	var rootNode ASTNode.ASTNode = c.prog(&tokens)
	return rootNode
}

func (c *SimpleCalculator) evaluate(node ASTNode.ASTNode, indent string) int {
	var result int = 0
	fmt.Println(indent + "Calculating: " + node.Typ)
	switch node.Typ {
	case ASTNodeType.Programm.String():
		for _, child := range node.GetChildren() {
			result = c.evaluate(child, indent+"\t")
		}
	case ASTNodeType.Additive.String():
		var child1 ASTNode.ASTNode = node.GetChild(0)
		var value1 int = c.evaluate(child1, indent+"\t")
		var child2 ASTNode.ASTNode = node.GetChild(1)
		var value2 int = c.evaluate(child2, indent+"\t")
		if node.Text == "+" {
			result = value1 + value2
		} else {
			result = value1 - value2
		}
	case ASTNodeType.Multiplicative.String():
		child1 := node.GetChild(0)
		value1 := c.evaluate(child1, indent+"\t")
		child2 := node.GetChild(1)
		value2 := c.evaluate(child2, indent+"\t")
		if node.Text == "*" {
			result = value1 * value2
		} else {
			result = value1 / value2
		}
	case ASTNodeType.IntLiteral.String():
		result, _ = strconv.Atoi(node.Text)
	default:

	}
	fmt.Println(indent + "Result: " + string(result))
	return result
}

func (c *SimpleCalculator) prog(tokens *lexer.SimpleTokenReader) ASTNode.ASTNode {
	var node ASTNode.ASTNode = ASTNode.ASTNode{Typ: ASTNodeType.Programm.String(), Text: "Calculator"}
	var child ASTNode.ASTNode = c.additive(tokens)

	if child.Text != "" {
		node.AddChild(child)
	}
	return node
}

func (c *SimpleCalculator) dumpAST(node ASTNode.ASTNode, indent string) {
	fmt.Println(indent + node.GetType() + " " + node.GetText())
	for _, child := range node.Children() {
		c.dumpAST(child, indent+"\t")
	}
}

/* 整型变量声明语句，如： */
func (c *SimpleCalculator) intDeclare(tokens *lexer.SimpleTokenReader) ASTNode.ASTNode {
	var node ASTNode.ASTNode = ASTNode.ASTNode{}
	var token lexer.SimpleToken = tokens.Peek()
	if token.Text != "" && token.Type == TokenType.Int.String() {
		token = tokens.Read()
		if tokens.Peek().Type == TokenType.Identifier.String() {
			token = tokens.Read()
			node = ASTNode.ASTNode{Typ: ASTNodeType.IntDeclaration.String(), Text: token.Text}
			token = tokens.Peek()
			if token.Text != "" && token.Type == TokenType.Assignment.String() {
				tokens.Read()
				var child ASTNode.ASTNode = c.additive(tokens)
				if child.Text == "" {
					fmt.Println("invalide variable initialization, expecting an expression")
					return ASTNode.ASTNode{}
				} else {
					node.AddChild(child)
				}
			}
		} else {
			fmt.Println("variable name expected")
		}
		if node.Text != "" {
			token = tokens.Peek()
			if token.Text != "" && token.Type == TokenType.SemiColon.String() {
				tokens.Read()
			} else {
				fmt.Println("invalid statement, expecting semicolon")
				return ASTNode.ASTNode{}
			}
		}
	}
	return node
}

/* 语法解析：加法表达式 */
func (c *SimpleCalculator) additive(tokens *lexer.SimpleTokenReader) ASTNode.ASTNode {
	var child1 ASTNode.ASTNode = c.multiplicative(tokens)
	var node ASTNode.ASTNode = child1

	var token lexer.SimpleToken = tokens.Peek()
	if child1.Text != "" && token.Text != "" {
		if token.Type == TokenType.Plus.String() || token.Type == TokenType.Minus.String() {
			token = tokens.Read()
			var child2 ASTNode.ASTNode = c.additive(tokens)
			if child2.Text != "" {
				node = ASTNode.ASTNode{Typ: ASTNodeType.Additive.String(), Text: token.Text}
				node.AddChild(child1)
				node.AddChild(child2)
			} else {
				fmt.Println("invalid additive expression, expecting the right part.")
				return ASTNode.ASTNode{}
			}
		}
	}
	return node
}

/* 乘法表达式 */
func (c *SimpleCalculator) multiplicative(tokens *lexer.SimpleTokenReader) ASTNode.ASTNode {
	var child1 ASTNode.ASTNode = c.primary(tokens)
	var node = child1

	var token lexer.SimpleToken = tokens.Peek()
	if child1.Text != "" && token.Text != "" {
		if token.Type == TokenType.Star.String() || token.Type == TokenType.Slash.String() {
			token = tokens.Read()
			var child2 ASTNode.ASTNode = c.primary(tokens)
			if child2.Text != "" {
				node = ASTNode.ASTNode{Typ: ASTNodeType.Multiplicative.String(), Text: token.Text}
				node.AddChild(child1)
				node.AddChild(child2)
			} else {
				fmt.Println("invalid multiplicative expression, expecting the right part.")
				return ASTNode.ASTNode{}
			}
		}
	}
	return node
}

/* 语法解析： 基础表达式 */
func (c *SimpleCalculator) primary(tokens *lexer.SimpleTokenReader) ASTNode.ASTNode {
	var node ASTNode.ASTNode
	var token lexer.SimpleToken = tokens.Peek()
	if token.Text != "" {
		if token.Type == TokenType.IntLiteral.String() {
			token = tokens.Read()
			node = ASTNode.ASTNode{Typ: ASTNodeType.IntLiteral.String(), Text: token.Text}
		} else if token.Type == TokenType.Identifier.String() {
			token = tokens.Read()
			node = ASTNode.ASTNode{Typ: ASTNodeType.Identifier.String(), Text: token.Text}
		} else if token.Type == TokenType.LeftParen.String() {
			tokens.Read()
			node = c.additive(tokens)
			if node.Text != "" {
				token = tokens.Peek()
				if token.Text != "" && token.Type == TokenType.RightParen.String() {
					tokens.Read()
				} else {
					fmt.Println("expecting right parent")
					return ASTNode.ASTNode{}
				}
			} else {
				fmt.Println("expecting an additive expression inside parenthesis")
			}
		}
	}
	// 这个方法也做了AST的简化，就是不用构造一个primary节点，直接返回子节点。因为它只有一个子节点。
	return node
}
