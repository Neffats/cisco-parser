package lexer

type Lexer struct {
	Name   string
	Input  string
	Tokens chan lexertoken.Token
	State  LexerFn

	Start int
	Pos   int
	Width int
}

func (this *Lexer) Backup() {
	this.Pos -= this.Width
}

func (this *Lexer) CurrentInput() string {
	return this.Input[this.Start:this.Pos]
}

func (this *Lexer) Dec() {
	this.Pos--
}

func (this *Lexer) Emit(tokenType lexerToken.TokenType) {
	this.Tokens <- lexerToken.Token{
		Type:  tokenType,
		Value: this.Input[this.Start:this.Pos],
	}
	this.Start = this.Pos
}
