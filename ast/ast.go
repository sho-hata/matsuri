package ast

import "github.com/sho-hata/matsuri/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// let <identifier> = <expression>;
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

var _ Statement = (*LetStatement)(nil)

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// return <expression>;

type ReturnStatement struct {
	Token       token.Token // 'return' token
	ReturnValue Expression
}

var _ Statement = (*ReturnStatement)(nil)

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}
