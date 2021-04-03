package parser

import (
	"fmt"
	"io"
	"strconv"
)

type bailout struct{}

var stmtStart = map[Token]bool{
	TokenBreak:    true,
	TokenContinue: true,
	TokenFor:      true,
	TokenIf:       true,
	TokenReturn:   true,
	TokenExport:   true,
}

// Parser parses the gslang source files. It's based on Go's parser
// implementation.
type Parser struct {
	file      *Code
	errors    ErrorList
	scanner   *Scanner
	pos       Pos
	token     Token
	tokenLit  string
	exprLevel int // < 0: in control clause, >= 0: in expression
	syncPos   Pos // last sync position
	syncCount int // number of advance calls without progress
	trace     bool
	indent    int
	traceOut  io.Writer
}

// NewParser creates a Parser.
func NewParser(file *Code, src []byte, trace io.Writer) *Parser {
	p := &Parser{
		file:     file,
		trace:    trace != nil,
		traceOut: trace,
	}
	p.scanner = NewScanner(p.file, src,
		func(pos FilePos, msg string) {
			p.errors.Add(pos, msg)
		}, 0)
	p.next()
	return p
}

// ParseFile parses the source and returns an AST file unit.
func (p *Parser) ParseFile() (file *File, err error) {
	defer func() {
		if e := recover(); e != nil {
			if _, ok := e.(bailout); !ok {
				panic(e)
			}
		}

		p.errors.Sort()
		err = p.errors.Err()
	}()

	if p.trace {
		defer untracep(tracep(p, "File"))
	}

	if p.errors.Len() > 0 {
		return nil, p.errors.Err()
	}

	stmts := p.parseStmtList()
	if p.errors.Len() > 0 {
		return nil, p.errors.Err()
	}

	file = &File{
		Input: p.file,
		Stmts: stmts,
	}
	return
}

func (p *Parser) parseExpr() Expr {
	if p.trace {
		defer untracep(tracep(p, "Expression"))
	}

	expr := p.parseBinaryExpr(LowestPrec + 1)

	// ternary conditional expression
	if p.token == TokenQuestion {
		return p.parseCondExpr(expr)
	}
	return expr
}

func (p *Parser) parseBinaryExpr(prec1 int) Expr {
	if p.trace {
		defer untracep(tracep(p, "BinaryExpression"))
	}

	x := p.parseUnaryExpr()

	for {
		op, prec := p.token, p.token.Precedence()
		if prec < prec1 {
			return x
		}

		pos := p.expect(op)

		y := p.parseBinaryExpr(prec + 1)

		x = &BinaryExpr{
			LHS:      x,
			RHS:      y,
			Token:    op,
			TokenPos: pos,
		}
	}
}

func (p *Parser) parseCondExpr(cond Expr) Expr {
	questionPos := p.expect(TokenQuestion)
	trueExpr := p.parseExpr()
	colonPos := p.expect(TokenColon)
	falseExpr := p.parseExpr()

	return &CondExpr{
		Cond:        cond,
		True:        trueExpr,
		False:       falseExpr,
		QuestionPos: questionPos,
		ColonPos:    colonPos,
	}
}

func (p *Parser) parseUnaryExpr() Expr {
	if p.trace {
		defer untracep(tracep(p, "UnaryExpression"))
	}

	switch p.token {
	case TokenAdd, TokenSub, TokenNot, TokenXor:
		pos, op := p.pos, p.token
		p.next()
		x := p.parseUnaryExpr()
		return &UnaryExpr{
			Token:    op,
			TokenPos: pos,
			Expr:     x,
		}
	}
	return p.parsePrimaryExpr()
}

func (p *Parser) parsePrimaryExpr() Expr {
	if p.trace {
		defer untracep(tracep(p, "PrimaryExpression"))
	}

	x := p.parseOperand()

L:
	for {
		switch p.token {
		case TokenPeriod:
			p.next()

			switch p.token {
			case TokenIdent:
				x = p.parseSelector(x)
			default:
				pos := p.pos
				p.errorExpected(pos, "selector")
				p.advance(stmtStart)
				return &BadExpr{From: pos, To: p.pos}
			}
		case TokenLBrack:
			x = p.parseIndexOrSlice(x)
		case TokenLParen:
			x = p.parseCall(x)
		default:
			break L
		}
	}
	return x
}

func (p *Parser) parseCall(x Expr) *CallExpr {
	if p.trace {
		defer untracep(tracep(p, "Call"))
	}

	lparen := p.expect(TokenLParen)
	p.exprLevel++

	var list []Expr
	var ellipsis Pos
	for p.token != TokenRParen && p.token != TokenEOF && !ellipsis.IsValid() {
		list = append(list, p.parseExpr())
		if p.token == TokenEllipsis {
			ellipsis = p.pos
			p.next()
		}
		if !p.expectComma(TokenRParen, "call argument") {
			break
		}
	}

	p.exprLevel--
	rparen := p.expect(TokenRParen)
	return &CallExpr{
		Func:     x,
		LParen:   lparen,
		RParen:   rparen,
		Ellipsis: ellipsis,
		Args:     list,
	}
}

func (p *Parser) expectComma(closing Token, want string) bool {
	if p.token == TokenComma {
		p.next()

		if p.token == closing {
			p.errorExpected(p.pos, want)
			return false
		}
		return true
	}

	if p.token == TokenSemicolon && p.tokenLit == "\n" {
		p.next()
	}
	return false
}

func (p *Parser) parseIndexOrSlice(x Expr) Expr {
	if p.trace {
		defer untracep(tracep(p, "IndexOrSlice"))
	}

	lbrack := p.expect(TokenLBrack)
	p.exprLevel++

	var index [2]Expr
	if p.token != TokenColon {
		index[0] = p.parseExpr()
	}
	numColons := 0
	if p.token == TokenColon {
		numColons++
		p.next()

		if p.token != TokenRBrack && p.token != TokenEOF {
			index[1] = p.parseExpr()
		}
	}

	p.exprLevel--
	rbrack := p.expect(TokenRBrack)

	if numColons > 0 {
		// slice expression
		return &SliceExpr{
			Expr:   x,
			LBrack: lbrack,
			RBrack: rbrack,
			Low:    index[0],
			High:   index[1],
		}
	}
	return &IndexExpr{
		Expr:   x,
		LBrack: lbrack,
		RBrack: rbrack,
		Index:  index[0],
	}
}

func (p *Parser) parseSelector(x Expr) Expr {
	if p.trace {
		defer untracep(tracep(p, "Selector"))
	}

	sel := p.parseIdent()
	return &SelectorExpr{Expr: x, Sel: &StringLit{
		Value:    sel.Name,
		ValuePos: sel.NamePos,
		Literal:  sel.Name,
	}}
}

func (p *Parser) parseOperand() Expr {
	if p.trace {
		defer untracep(tracep(p, "Operand"))
	}

	switch p.token {
	case TokenIdent:
		return p.parseIdent()
	case TokenInt:
		v, _ := strconv.ParseInt(p.tokenLit, 10, 64)
		x := &IntLit{
			Value:    v,
			ValuePos: p.pos,
			Literal:  p.tokenLit,
		}
		p.next()
		return x
	case TokenFloat:
		v, _ := strconv.ParseFloat(p.tokenLit, 64)
		x := &FloatLit{
			Value:    v,
			ValuePos: p.pos,
			Literal:  p.tokenLit,
		}
		p.next()
		return x
	case TokenChar:
		return p.parseCharLit()
	case TokenString:
		v, _ := strconv.Unquote(p.tokenLit)
		x := &StringLit{
			Value:    v,
			ValuePos: p.pos,
			Literal:  p.tokenLit,
		}
		p.next()
		return x
	case TokenTrue:
		x := &BoolLit{
			Value:    true,
			ValuePos: p.pos,
			Literal:  p.tokenLit,
		}
		p.next()
		return x
	case TokenFalse:
		x := &BoolLit{
			Value:    false,
			ValuePos: p.pos,
			Literal:  p.tokenLit,
		}
		p.next()
		return x
	case TokenNil:
		x := &NilLit{TokenPos: p.pos}
		p.next()
		return x
	case TokenImport:
		return p.parseImportExpr()
	case TokenLParen:
		lparen := p.pos
		p.next()
		p.exprLevel++
		x := p.parseExpr()
		p.exprLevel--
		rparen := p.expect(TokenRParen)
		return &ParenExpr{
			LParen: lparen,
			Expr:   x,
			RParen: rparen,
		}
	case TokenLBrack: // array literal
		return p.parseArrayLit()
	case TokenLBrace: // map literal
		return p.parseMapLit()
	case TokenFunc: // function literal
		return p.parseFuncLit()
	case TokenError: // error expression
		return p.parseErrorExpr()
	}

	pos := p.pos
	p.errorExpected(pos, "operand")
	p.advance(stmtStart)
	return &BadExpr{From: pos, To: p.pos}
}

func (p *Parser) parseImportExpr() Expr {
	pos := p.pos
	p.next()
	p.expect(TokenLParen)
	if p.token != TokenString {
		p.errorExpected(p.pos, "module name")
		p.advance(stmtStart)
		return &BadExpr{From: pos, To: p.pos}
	}

	// module name
	moduleName, _ := strconv.Unquote(p.tokenLit)
	expr := &ImportExpr{
		ModuleName: moduleName,
		Token:      TokenImport,
		TokenPos:   pos,
	}

	p.next()
	p.expect(TokenRParen)
	return expr
}

func (p *Parser) parseCharLit() Expr {
	if n := len(p.tokenLit); n >= 3 {
		code, _, _, err := strconv.UnquoteChar(p.tokenLit[1:n-1], '\'')
		if err == nil {
			x := &CharLit{
				Value:    code,
				ValuePos: p.pos,
				Literal:  p.tokenLit,
			}
			p.next()
			return x
		}
	}

	pos := p.pos
	p.error(pos, "illegal char literal")
	p.next()
	return &BadExpr{
		From: pos,
		To:   p.pos,
	}
}

func (p *Parser) parseFuncLit() Expr {
	if p.trace {
		defer untracep(tracep(p, "FuncLit"))
	}

	typ := p.parseFuncType()
	p.exprLevel++
	body := p.parseBody()
	p.exprLevel--
	return &FuncLit{
		Type: typ,
		Body: body,
	}
}

func (p *Parser) parseArrayLit() Expr {
	if p.trace {
		defer untracep(tracep(p, "ArrayLit"))
	}

	lbrack := p.expect(TokenLBrack)
	p.exprLevel++

	var elements []Expr
	for p.token != TokenRBrack && p.token != TokenEOF {
		elements = append(elements, p.parseExpr())

		if !p.expectComma(TokenRBrack, "array element") {
			break
		}
	}

	p.exprLevel--
	rbrack := p.expect(TokenRBrack)
	return &ArrayLit{
		Elements: elements,
		LBrack:   lbrack,
		RBrack:   rbrack,
	}
}

func (p *Parser) parseErrorExpr() Expr {
	pos := p.pos

	p.next()
	lparen := p.expect(TokenLParen)
	value := p.parseExpr()
	rparen := p.expect(TokenRParen)
	return &ErrorExpr{
		ErrorPos: pos,
		Expr:     value,
		LParen:   lparen,
		RParen:   rparen,
	}
}

func (p *Parser) parseFuncType() *FuncType {
	if p.trace {
		defer untracep(tracep(p, "FuncType"))
	}

	pos := p.expect(TokenFunc)
	params := p.parseIdentList()
	return &FuncType{
		FuncPos: pos,
		Params:  params,
	}
}

func (p *Parser) parseBody() *BlockStmt {
	if p.trace {
		defer untracep(tracep(p, "Body"))
	}

	lbrace := p.expect(TokenLBrace)
	list := p.parseStmtList()
	rbrace := p.expect(TokenRBrace)
	return &BlockStmt{
		LBrace: lbrace,
		RBrace: rbrace,
		Stmts:  list,
	}
}

func (p *Parser) parseStmtList() (list []Stmt) {
	if p.trace {
		defer untracep(tracep(p, "StatementList"))
	}

	for p.token != TokenRBrace && p.token != TokenEOF {
		list = append(list, p.parseStmt())
	}
	return
}

func (p *Parser) parseIdent() *Ident {
	pos := p.pos
	name := "_"

	if p.token == TokenIdent {
		name = p.tokenLit
		p.next()
	} else {
		p.expect(TokenIdent)
	}
	return &Ident{
		NamePos: pos,
		Name:    name,
	}
}

func (p *Parser) parseIdentList() *IdentList {
	if p.trace {
		defer untracep(tracep(p, "IdentList"))
	}

	var params []*Ident
	lparen := p.expect(TokenLParen)
	isVarArgs := false
	if p.token != TokenRParen {
		if p.token == TokenEllipsis {
			isVarArgs = true
			p.next()
		}

		params = append(params, p.parseIdent())
		for !isVarArgs && p.token == TokenComma {
			p.next()
			if p.token == TokenEllipsis {
				isVarArgs = true
				p.next()
			}
			params = append(params, p.parseIdent())
		}
	}

	rparen := p.expect(TokenRParen)
	return &IdentList{
		LParen:  lparen,
		RParen:  rparen,
		VarArgs: isVarArgs,
		List:    params,
	}
}

func (p *Parser) parseStmt() (stmt Stmt) {
	if p.trace {
		defer untracep(tracep(p, "Statement"))
	}

	switch p.token {
	case // simple statements
		TokenFunc, TokenError, TokenIdent, TokenInt,
		TokenFloat, TokenChar, TokenString, TokenTrue, TokenFalse,
		TokenNil, TokenImport, TokenLParen, TokenLBrace,
		TokenLBrack, TokenAdd, TokenSub, TokenMul, TokenAnd, TokenXor,
		TokenNot:
		s := p.parseSimpleStmt(false)
		p.expectSemi()
		return s
	case TokenReturn:
		return p.parseReturnStmt()
	case TokenExport:
		return p.parseExportStmt()
	case TokenIf:
		return p.parseIfStmt()
	case TokenFor:
		return p.parseForStmt()
	case TokenBreak, TokenContinue:
		return p.parseBranchStmt(p.token)
	case TokenSemicolon:
		s := &EmptyStmt{Semicolon: p.pos, Implicit: p.tokenLit == "\n"}
		p.next()
		return s
	case TokenRBrace:
		// semicolon may be omitted before a closing "}"
		return &EmptyStmt{Semicolon: p.pos, Implicit: true}
	default:
		pos := p.pos
		p.errorExpected(pos, "statement")
		p.advance(stmtStart)
		return &BadStmt{From: pos, To: p.pos}
	}
}

func (p *Parser) parseForStmt() Stmt {
	if p.trace {
		defer untracep(tracep(p, "ForStmt"))
	}

	pos := p.expect(TokenFor)

	// for {}
	if p.token == TokenLBrace {
		body := p.parseBlockStmt()
		p.expectSemi()

		return &ForStmt{
			ForPos: pos,
			Body:   body,
		}
	}

	prevLevel := p.exprLevel
	p.exprLevel = -1

	var s1 Stmt
	if p.token != TokenSemicolon { // skipping init
		s1 = p.parseSimpleStmt(true)
	}

	// for _ in seq {}            or
	// for value in seq {}        or
	// for key, value in seq {}
	if forInStmt, isForIn := s1.(*ForInStmt); isForIn {
		forInStmt.ForPos = pos
		p.exprLevel = prevLevel
		forInStmt.Body = p.parseBlockStmt()
		p.expectSemi()
		return forInStmt
	}

	// for init; cond; post {}
	var s2, s3 Stmt
	if p.token == TokenSemicolon {
		p.next()
		if p.token != TokenSemicolon {
			s2 = p.parseSimpleStmt(false) // cond
		}
		p.expect(TokenSemicolon)
		if p.token != TokenLBrace {
			s3 = p.parseSimpleStmt(false) // post
		}
	} else {
		// for cond {}
		s2 = s1
		s1 = nil
	}

	// body
	p.exprLevel = prevLevel
	body := p.parseBlockStmt()
	p.expectSemi()
	cond := p.makeExpr(s2, "condition expression")
	return &ForStmt{
		ForPos: pos,
		Init:   s1,
		Cond:   cond,
		Post:   s3,
		Body:   body,
	}
}

func (p *Parser) parseBranchStmt(tok Token) Stmt {
	if p.trace {
		defer untracep(tracep(p, "BranchStmt"))
	}

	pos := p.expect(tok)

	var label *Ident
	if p.token == TokenIdent {
		label = p.parseIdent()
	}
	p.expectSemi()
	return &BranchStmt{
		Token:    tok,
		TokenPos: pos,
		Label:    label,
	}
}

func (p *Parser) parseIfStmt() Stmt {
	if p.trace {
		defer untracep(tracep(p, "IfStmt"))
	}

	pos := p.expect(TokenIf)
	init, cond := p.parseIfHeader()
	body := p.parseBlockStmt()

	var elseStmt Stmt
	if p.token == TokenElse {
		p.next()

		switch p.token {
		case TokenIf:
			elseStmt = p.parseIfStmt()
		case TokenLBrace:
			elseStmt = p.parseBlockStmt()
			p.expectSemi()
		default:
			p.errorExpected(p.pos, "if or {")
			elseStmt = &BadStmt{From: p.pos, To: p.pos}
		}
	} else {
		p.expectSemi()
	}
	return &IfStmt{
		IfPos: pos,
		Init:  init,
		Cond:  cond,
		Body:  body,
		Else:  elseStmt,
	}
}

func (p *Parser) parseBlockStmt() *BlockStmt {
	if p.trace {
		defer untracep(tracep(p, "BlockStmt"))
	}

	lbrace := p.expect(TokenLBrace)
	list := p.parseStmtList()
	rbrace := p.expect(TokenRBrace)
	return &BlockStmt{
		LBrace: lbrace,
		RBrace: rbrace,
		Stmts:  list,
	}
}

func (p *Parser) parseIfHeader() (init Stmt, cond Expr) {
	if p.token == TokenLBrace {
		p.error(p.pos, "missing condition in if statement")
		cond = &BadExpr{From: p.pos, To: p.pos}
		return
	}

	outer := p.exprLevel
	p.exprLevel = -1
	if p.token == TokenSemicolon {
		p.error(p.pos, "missing init in if statement")
		return
	}
	init = p.parseSimpleStmt(false)

	var condStmt Stmt
	if p.token == TokenLBrace {
		condStmt = init
		init = nil
	} else if p.token == TokenSemicolon {
		p.next()

		condStmt = p.parseSimpleStmt(false)
	} else {
		p.error(p.pos, "missing condition in if statement")
	}

	if condStmt != nil {
		cond = p.makeExpr(condStmt, "boolean expression")
	}
	if cond == nil {
		cond = &BadExpr{From: p.pos, To: p.pos}
	}
	p.exprLevel = outer
	return
}

func (p *Parser) makeExpr(s Stmt, want string) Expr {
	if s == nil {
		return nil
	}

	if es, isExpr := s.(*ExprStmt); isExpr {
		return es.Expr
	}

	found := "simple statement"
	if _, isAss := s.(*AssignStmt); isAss {
		found = "assignment"
	}
	p.error(s.Pos(), fmt.Sprintf("expected %s, found %s", want, found))
	return &BadExpr{From: s.Pos(), To: p.safePos(s.End())}
}

func (p *Parser) parseReturnStmt() Stmt {
	if p.trace {
		defer untracep(tracep(p, "ReturnStmt"))
	}

	pos := p.pos
	p.expect(TokenReturn)

	var x Expr
	if p.token != TokenSemicolon && p.token != TokenRBrace {
		x = p.parseExpr()
	}
	p.expectSemi()
	return &ReturnStmt{
		ReturnPos: pos,
		Result:    x,
	}
}

func (p *Parser) parseExportStmt() Stmt {
	if p.trace {
		defer untracep(tracep(p, "ExportStmt"))
	}

	pos := p.pos
	p.expect(TokenExport)
	x := p.parseExpr()
	p.expectSemi()
	return &ExportStmt{
		ExportPos: pos,
		Result:    x,
	}
}

func (p *Parser) parseSimpleStmt(forIn bool) Stmt {
	if p.trace {
		defer untracep(tracep(p, "SimpleStmt"))
	}

	x := p.parseExprList()

	switch p.token {
	case TokenAssign, TokenDefine: // assignment statement
		pos, tok := p.pos, p.token
		p.next()
		y := p.parseExprList()
		return &AssignStmt{
			LHS:      x,
			RHS:      y,
			Token:    tok,
			TokenPos: pos,
		}
	case TokenIn:
		if forIn {
			p.next()
			y := p.parseExpr()

			var key, value *Ident
			var ok bool
			switch len(x) {
			case 1:
				key = &Ident{Name: "_", NamePos: x[0].Pos()}

				value, ok = x[0].(*Ident)
				if !ok {
					p.errorExpected(x[0].Pos(), "identifier")
					value = &Ident{Name: "_", NamePos: x[0].Pos()}
				}
			case 2:
				key, ok = x[0].(*Ident)
				if !ok {
					p.errorExpected(x[0].Pos(), "identifier")
					key = &Ident{Name: "_", NamePos: x[0].Pos()}
				}
				value, ok = x[1].(*Ident)
				if !ok {
					p.errorExpected(x[1].Pos(), "identifier")
					value = &Ident{Name: "_", NamePos: x[1].Pos()}
				}
			}
			return &ForInStmt{
				Key:      key,
				Value:    value,
				Iterable: y,
			}
		}
	}

	if len(x) > 1 {
		p.errorExpected(x[0].Pos(), "1 expression")
		// continue with first expression
	}

	switch p.token {
	case TokenDefine,
		TokenAddAssign, TokenSubAssign, TokenMulAssign, TokenQuoAssign,
		TokenRemAssign, TokenAndAssign, TokenOrAssign, TokenXorAssign,
		TokenShlAssign, TokenShrAssign, TokenAndNotAssign:
		pos, tok := p.pos, p.token
		p.next()
		y := p.parseExpr()
		return &AssignStmt{
			LHS:      []Expr{x[0]},
			RHS:      []Expr{y},
			Token:    tok,
			TokenPos: pos,
		}
	case TokenInc, TokenDec:
		// increment or decrement statement
		s := &IncDecStmt{Expr: x[0], Token: p.token, TokenPos: p.pos}
		p.next()
		return s
	}
	return &ExprStmt{Expr: x[0]}
}

func (p *Parser) parseExprList() (list []Expr) {
	if p.trace {
		defer untracep(tracep(p, "ExpressionList"))
	}

	list = append(list, p.parseExpr())
	for p.token == TokenComma {
		p.next()
		list = append(list, p.parseExpr())
	}
	return
}

func (p *Parser) parseMapElementLit() *MapElementLit {
	if p.trace {
		defer untracep(tracep(p, "MapElementLit"))
	}

	pos := p.pos
	name := "_"
	if p.token == TokenIdent {
		name = p.tokenLit
	} else if p.token == TokenString {
		v, _ := strconv.Unquote(p.tokenLit)
		name = v
	} else {
		p.errorExpected(pos, "map key")
	}
	p.next()
	colonPos := p.expect(TokenColon)
	valueExpr := p.parseExpr()
	return &MapElementLit{
		Key:      name,
		KeyPos:   pos,
		ColonPos: colonPos,
		Value:    valueExpr,
	}
}

func (p *Parser) parseMapLit() *MapLit {
	if p.trace {
		defer untracep(tracep(p, "MapLit"))
	}

	lbrace := p.expect(TokenLBrace)
	p.exprLevel++

	var elements []*MapElementLit
	for p.token != TokenRBrace && p.token != TokenEOF {
		elements = append(elements, p.parseMapElementLit())

		if !p.expectComma(TokenRBrace, "map element") {
			break
		}
	}

	p.exprLevel--
	rbrace := p.expect(TokenRBrace)
	return &MapLit{
		LBrace:   lbrace,
		RBrace:   rbrace,
		Elements: elements,
	}
}

func (p *Parser) expect(token Token) Pos {
	pos := p.pos

	if p.token != token {
		p.errorExpected(pos, "'"+token.String()+"'")
	}
	p.next()
	return pos
}

func (p *Parser) expectSemi() {
	switch p.token {
	case TokenRParen, TokenRBrace:
		// semicolon is optional before a closing ')' or '}'
	case TokenComma:
		// permit a ',' instead of a ';' but complain
		p.errorExpected(p.pos, "';'")
		fallthrough
	case TokenSemicolon:
		p.next()
	default:
		p.errorExpected(p.pos, "';'")
		p.advance(stmtStart)
	}
}

func (p *Parser) advance(to map[Token]bool) {
	for ; p.token != TokenEOF; p.next() {
		if to[p.token] {
			if p.pos == p.syncPos && p.syncCount < 10 {
				p.syncCount++
				return
			}
			if p.pos > p.syncPos {
				p.syncPos = p.pos
				p.syncCount = 0
				return
			}
		}
	}
}

func (p *Parser) error(pos Pos, msg string) {
	filePos := p.file.Position(pos)

	n := len(p.errors)
	if n > 0 && p.errors[n-1].Pos.Line == filePos.Line {
		// discard errors reported on the same line
		return
	}
	if n > 10 {
		// too many errors; terminate early
		panic(bailout{})
	}
	p.errors.Add(filePos, msg)
}

func (p *Parser) errorExpected(pos Pos, msg string) {
	msg = "expected " + msg
	if pos == p.pos {
		// error happened at the current position: provide more specific
		switch {
		case p.token == TokenSemicolon && p.tokenLit == "\n":
			msg += ", found newline"
		case p.token.IsLiteral():
			msg += ", found " + p.tokenLit
		default:
			msg += ", found '" + p.token.String() + "'"
		}
	}
	p.error(pos, msg)
}

func (p *Parser) next() {
	if p.trace && p.pos.IsValid() {
		s := p.token.String()
		switch {
		case p.token.IsLiteral():
			p.printTrace(s, p.tokenLit)
		case p.token.IsOperator(), p.token.IsKeyword():
			p.printTrace(`"` + s + `"`)
		default:
			p.printTrace(s)
		}
	}
	p.token, p.tokenLit, p.pos = p.scanner.Scan()
}

func (p *Parser) printTrace(a ...interface{}) {
	const (
		dots = ". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . "
		n    = len(dots)
	)

	filePos := p.file.Position(p.pos)
	_, _ = fmt.Fprintf(p.traceOut, "%5d: %5d:%3d: ", p.pos, filePos.Line,
		filePos.Column)
	i := 2 * p.indent
	for i > n {
		_, _ = fmt.Fprint(p.traceOut, dots)
		i -= n
	}
	_, _ = fmt.Fprint(p.traceOut, dots[0:i])
	_, _ = fmt.Fprintln(p.traceOut, a...)
}

func (p *Parser) safePos(pos Pos) Pos {
	fileBase := p.file.Base
	fileSize := p.file.Size

	if int(pos) < fileBase || int(pos) > fileBase+fileSize {
		return Pos(fileBase + fileSize)
	}
	return pos
}

func tracep(p *Parser, msg string) *Parser {
	p.printTrace(msg, "(")
	p.indent++
	return p
}

func untracep(p *Parser) {
	p.indent--
	p.printTrace(")")
}
