package app_context

type AppContext interface {
}

type appContext struct {
}

func NewAppContext() *appContext {
	return &appContext{}
}
