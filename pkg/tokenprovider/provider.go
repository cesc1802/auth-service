package tokenprovider

type Provider interface {
	Generate(data TokenPayload) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}
