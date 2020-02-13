package lexer

import (
	"strings"
	"unicode"

	"github.com/Neffats/cisco-parser/services/lexer/lexertoken"
)

func LexAccessList(l *Lexer) LexFn {
	l.SkipWhitespace()
	for !unicode.IsSpace(l.Next()) {
	}

	l.Dec()
	l.Emit(lexertoken.TOKEN_ACL_NAME)

	if strings.HasPrefix(l.Input[l.Pos:], lexertoken.Line) {
		return LexLine
	}

	return LexExtended
}
