package lexer

import (
	"strings"

	"github.com/Neffats/cisco-parser/services/lexer/lexertoken"
)

func LexBegin(l *Lexer) LexFn {
	for {
		if strings.HasPrefix(l.Input[l.Pos:], lexertoken.ACL) {
			if l.Pos > l.Start {
				l.Emit(lexertoken.TOKEN_ACL)
				return LexAccessList
			}
		}

		if l.Next() == lexertoken.EOF {
			break
		}
	}

	if l.Pos > l.Start {
		l.Emit(lexertoken.TOKEN_ACL)
	}
	l.Emit(lexertoken.TOKEN_EOF)
	return nil
}
