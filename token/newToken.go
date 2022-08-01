package token

type Token struct {
	Name string
}

func NewToken(token *Token) string {
	return "一串数字" + token.Name
}
