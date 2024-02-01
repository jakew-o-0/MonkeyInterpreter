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
