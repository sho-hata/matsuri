package token

import "testing"

func TestLookupIdent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		ident string
		want  TokenType
	}{
		{"function", "fn", FUNCTION},
		{"let", "let", LET},
		{"hoge", "hoge", IDENT},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := LookupIdent(tt.ident); got != tt.want {
				t.Errorf("LookupIdent() = %v, want %v", got, tt.want)
			}
		})
	}
}
