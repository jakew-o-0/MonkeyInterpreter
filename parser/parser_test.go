package parser

import (
    "testing"
    "github.com/jakew-o-0/MonkeyInterpreter/ast"
    "github.com/jakew-o-0/MonkeyInterpreter/lexer"
)


func TestLetStatements(t *testing.T) {
    input := `
    let x = 5;
    let y = 10;
    let foobar = 69420;
    `

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)

    if program == nil {
	t.Fatalf("ParseProgram() returned nil")
    }
    if len(program.Statements) != 3 {
	t.Fatalf("program does not contain 3 statements. got=%d", len(program.Statements))
    } 
    
    tests := []struct {
	expectedIdentifier string
    }{
	{"x"},
	{"y"},
	{"foobar"},
    }

    for i, tt := range tests {
	stmt := program.Statements[i]
	if !testLetStatements(t, stmt, tt.expectedIdentifier) {
	    return
	}
    }
}

func testLetStatements(t *testing.T, s ast.Statement, name string) bool {
    if s.TokenLiteral() != "let" {
	t.Errorf("s.TokenLIteral not 'let'. got=%q", s.TokenLiteral())
	return false
    }
    
    letStmt, ok := s.(*ast.LetStatement)
    if !ok {
	t.Errorf("s not *ast.LetStatement. got=%T", s)
	return false
    }

    if letStmt.Name.Value != name {
	t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
	return false
    }

    return true
}

func checkParserErrors(t *testing.T, p *Parser) {
    errs := p.Errors()
    if len(errs) == 0 {
	return
    }

    t.Errorf("parser has %d errors", len(errs))
    for _, e := range errs {
	t.Errorf("parser Error: %q", e)
    }

    t.FailNow()
}







func TestReturnStatement(t *testing.T) {
    input := `
    return 5;
    return 10;
    return 69420;
    `

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t,p)

    if len(program.Statements) != 3 {
	t.Fatalf("program.statements does not contain 3 statements. got=%d", len(program.Statements))
    }

    for _,stmt := range program.Statements {
	returnstmt, ok :=  stmt.(*ast.ReturnStatement)
	if !ok {
	    t.Fatalf("stmt not *ast.ReturnStatement")
	}
	if returnstmt.TokenLiteral() != "return" {
	    t.Errorf("returnstmt.tokenLiteral not 'return'. got=%q", returnstmt.TokenLiteral())
	}
    }
}
