package lexertoken

import "fmt"

type Token struct {
	Type  TokenType
	Value string
}
