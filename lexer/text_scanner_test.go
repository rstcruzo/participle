package lexer

import (
	"strings"
	"testing"
	"text/scanner"

	"github.com/stretchr/testify/require"
)

func TestLexer(t *testing.T) {
	lexer := Upgrade(LexString("hello world"))
	helloPos := Position{Offset: 0, Line: 1, Column: 1}
	worldPos := Position{Offset: 6, Line: 1, Column: 7}
	eofPos := Position{Offset: 11, Line: 1, Column: 12}
	require.Equal(t, Token{Type: scanner.Ident, Value: "hello", Pos: helloPos}, mustPeek(t, lexer, 0))
	require.Equal(t, Token{Type: scanner.Ident, Value: "hello", Pos: helloPos}, mustPeek(t, lexer, 0))
	require.Equal(t, Token{Type: scanner.Ident, Value: "hello", Pos: helloPos}, mustNext(t, lexer))
	require.Equal(t, Token{Type: scanner.Ident, Value: "world", Pos: worldPos}, mustPeek(t, lexer, 0))
	require.Equal(t, Token{Type: scanner.Ident, Value: "world", Pos: worldPos}, mustNext(t, lexer))
	require.Equal(t, Token{Type: scanner.EOF, Value: "", Pos: eofPos}, mustPeek(t, lexer, 0))
	require.Equal(t, Token{Type: scanner.EOF, Value: "", Pos: eofPos}, mustNext(t, lexer))
}

func TestLexString(t *testing.T) {
	lexer := LexString(`"hello\nworld"`)
	token, err := lexer.Next()
	require.NoError(t, err)
	require.Equal(t, Token{Type: scanner.String, Value: "hello\nworld", Pos: Position{Line: 1, Column: 1}}, token)
}

func TestLexSingleString(t *testing.T) {
	lexer := LexString(`'hello\nworld'`)
	token, err := lexer.Next()
	require.NoError(t, err)
	require.Equal(t, Token{Type: scanner.String, Value: "hello\nworld", Pos: Position{Line: 1, Column: 1}}, token)
	lexer = LexString(`'\U00008a9e'`)
	token, err = lexer.Next()
	require.NoError(t, err)
	require.Equal(t, Token{Type: scanner.Char, Value: "\U00008a9e", Pos: Position{Line: 1, Column: 1}}, token)
}

func BenchmarkTextScannerLexer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lex, err := TextScannerLexer.Lex(strings.NewReader("hello world 123 hello world 123"))
		require.NoError(b, err)
		ConsumeAll(lex) // nolint: errcheck
	}
}
