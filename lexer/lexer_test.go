package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		expectedType    tokenType
		expectedLiteral string
	}{
		{INC_VALUE, "+"},
		{INC_POINTER, ">"},
		{DEC_VALUE, "-"},
		{DEC_POINTER, "<"},
		{LOOP_START, "["},
		{LOOP_END, "]"},
		{INPUT, ","},
		{OUTPUT, "."},
	}
	l := New("+>-<[],.")
	tokens := l.GetTokens()
	if len(tokens) != len(tests) {
		t.Fatalf("Expected %d length. Got %d", len(tests), len(tokens))
	}
	for index, token := range tokens {
		if token.Type != tests[index].expectedType {
			t.Fatalf("Expected %s type. Got %s", tests[index].expectedType, token.Type)
		}
		if token.LiteralValue != tests[index].expectedLiteral {
			t.Fatalf("dfred")
		}
	}
}
