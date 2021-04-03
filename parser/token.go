package parser

import "strconv"

var keywords map[string]Token

// Token represents a token.
type Token int

// List of tokens
const (
	TokenIllegal Token = iota
	TokenEOF
	TokenComment
	Token_literalBeg
	TokenIdent
	TokenInt
	TokenFloat
	TokenChar
	TokenString
	Token_literalEnd
	Token_operatorBeg
	TokenAdd          // +
	TokenSub          // -
	TokenMul          // *
	TokenQuo          // /
	TokenRem          // %
	TokenAnd          // &
	TokenOr           // |
	TokenXor          // ^
	TokenShl          // <<
	TokenShr          // >>
	TokenAndNot       // &^
	TokenAddAssign    // +=
	TokenSubAssign    // -=
	TokenMulAssign    // *=
	TokenQuoAssign    // /=
	TokenRemAssign    // %=
	TokenAndAssign    // &=
	TokenOrAssign     // |=
	TokenXorAssign    // ^=
	TokenShlAssign    // <<=
	TokenShrAssign    // >>=
	TokenAndNotAssign // &^=
	TokenLAnd         // &&
	TokenLOr          // ||
	TokenInc          // ++
	TokenDec          // --
	TokenEqual        // ==
	TokenLess         // <
	TokenGreater      // >
	TokenAssign       // =
	TokenNot          // !
	TokenNotEqual     // !=
	TokenLessEq       // <=
	TokenGreaterEq    // >=
	TokenDefine       // :=
	TokenEllipsis     // ...
	TokenLParen       // (
	TokenLBrack       // [
	TokenLBrace       // {
	TokenComma        // ,
	TokenPeriod       // .
	TokenRParen       // )
	TokenRBrack       // ]
	TokenRBrace       // }
	TokenSemicolon    // ;
	TokenColon        // :
	TokenQuestion     // ?
	Token_operatorEnd
	Token_keywordBeg
	TokenBreak
	TokenContinue
	TokenElse
	TokenFor
	TokenFunc
	TokenError
	TokenIf
	TokenReturn
	TokenExport
	TokenTrue
	TokenFalse
	TokenIn
	TokenNil
	TokenImport
	Token_keywordEnd
)

var tokens = [...]string{
	TokenIllegal:      "ILLEGAL",
	TokenEOF:          "EOF",
	TokenComment:      "COMMENT",
	TokenIdent:        "IDENT",
	TokenInt:          "INT",
	TokenFloat:        "FLOAT",
	TokenChar:         "CHAR",
	TokenString:       "STRING",
	TokenAdd:          "+",
	TokenSub:          "-",
	TokenMul:          "*",
	TokenQuo:          "/",
	TokenRem:          "%",
	TokenAnd:          "&",
	TokenOr:           "|",
	TokenXor:          "^",
	TokenShl:          "<<",
	TokenShr:          ">>",
	TokenAndNot:       "&^",
	TokenAddAssign:    "+=",
	TokenSubAssign:    "-=",
	TokenMulAssign:    "*=",
	TokenQuoAssign:    "/=",
	TokenRemAssign:    "%=",
	TokenAndAssign:    "&=",
	TokenOrAssign:     "|=",
	TokenXorAssign:    "^=",
	TokenShlAssign:    "<<=",
	TokenShrAssign:    ">>=",
	TokenAndNotAssign: "&^=",
	TokenLAnd:         "&&",
	TokenLOr:          "||",
	TokenInc:          "++",
	TokenDec:          "--",
	TokenEqual:        "==",
	TokenLess:         "<",
	TokenGreater:      ">",
	TokenAssign:       "=",
	TokenNot:          "!",
	TokenNotEqual:     "!=",
	TokenLessEq:       "<=",
	TokenGreaterEq:    ">=",
	TokenDefine:       ":=",
	TokenEllipsis:     "...",
	TokenLParen:       "(",
	TokenLBrack:       "[",
	TokenLBrace:       "{",
	TokenComma:        ",",
	TokenPeriod:       ".",
	TokenRParen:       ")",
	TokenRBrack:       "]",
	TokenRBrace:       "}",
	TokenSemicolon:    ";",
	TokenColon:        ":",
	TokenQuestion:     "?",
	TokenBreak:        "break",
	TokenContinue:     "continue",
	TokenElse:         "else",
	TokenFor:          "for",
	TokenFunc:         "func",
	TokenError:        "error",
	TokenIf:           "if",
	TokenReturn:       "return",
	TokenExport:       "export",
	TokenTrue:         "true",
	TokenFalse:        "false",
	TokenIn:           "in",
	TokenNil:    	   "nil",
	TokenImport:       "import",
}

func (tok Token) String() string {
	s := ""

	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}

	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}

	return s
}

// LowestPrec represents lowest operator precedence.
const LowestPrec = 0

// Precedence returns the precedence for the operator token.
func (tok Token) Precedence() int {
	switch tok {
	case TokenLOr:
		return 1
	case TokenLAnd:
		return 2
	case TokenEqual, TokenNotEqual, TokenLess, TokenLessEq, TokenGreater, TokenGreaterEq:
		return 3
	case TokenAdd, TokenSub, TokenOr, TokenXor:
		return 4
	case TokenMul, TokenQuo, TokenRem, TokenShl, TokenShr, TokenAnd, TokenAndNot:
		return 5
	}
	return LowestPrec
}

// IsLiteral returns true if the token is a literal.
func (tok Token) IsLiteral() bool {
	return Token_literalBeg < tok && tok < Token_literalEnd
}

// IsOperator returns true if the token is an operator.
func (tok Token) IsOperator() bool {
	return Token_operatorBeg < tok && tok < Token_operatorEnd
}

// IsKeyword returns true if the token is a keyword.
func (tok Token) IsKeyword() bool {
	return Token_keywordBeg < tok && tok < Token_keywordEnd
}

// Lookup returns corresponding keyword if ident is a keyword.
func Lookup(ident string) Token {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	return TokenIdent
}

func init() {
	keywords = make(map[string]Token)
	for i := Token_keywordBeg + 1; i < Token_keywordEnd; i++ {
		keywords[tokens[i]] = i
	}
}
