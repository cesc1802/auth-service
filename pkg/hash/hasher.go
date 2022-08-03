package hash

type Hasher interface {
	Hash(data string) string
}
