package ifaces

type Requester interface {
	GetUserID() uint
	GetRefreshTokenID() string
}
