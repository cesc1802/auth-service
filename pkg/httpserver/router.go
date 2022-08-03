package httpserver

import (
	"context"
	"fmt"
	"github.com/cesc1802/auth-service/pkg/httpserver/middleware"
	"github.com/cesc1802/auth-service/pkg/i18n"
	"github.com/cesc1802/auth-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strings"
)

type MyHttpServerConfig struct {
	mode string
	host string
	port string
}

func NewMyHttpServerConfig(mode, host, port string) *MyHttpServerConfig {
	return &MyHttpServerConfig{
		mode: mode,
		host: host,
		port: port,
	}
}

type myHttpServer struct {
	cnf       *MyHttpServerConfig
	engine    *gin.Engine
	isRunning bool
	*http.Server
	handlers []func(engine *gin.Engine)
	i18n     *i18n.AppI18n
	logger   *logger.Logger
}

func New(cnf *MyHttpServerConfig, i18n *i18n.AppI18n) *myHttpServer {
	return &myHttpServer{
		engine:   gin.New(),
		i18n:     i18n,
		cnf:      cnf,
		handlers: []func(*gin.Engine){},
		logger:   logger.New("gin-service"),
	}
}

func jsonTagNameFunc(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}
func (s *myHttpServer) Setup() error {
	if s.isRunning {
		return nil
	}
	if s.cnf.mode == "" {
		s.cnf.mode = "debug"
	}

	gin.SetMode(s.cnf.mode)
	s.engine = gin.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(jsonTagNameFunc)
	}

	s.engine.RedirectTrailingSlash = true
	s.engine.RedirectFixedPath = true

	// Recovery
	//TODO: you can add more middleware here
	s.engine.Use(middleware.Recovery(s.i18n))

	s.isRunning = true
	return nil
}

func (s *myHttpServer) Start() error {
	if err := s.Setup(); err != nil {
		return err
	}

	//TODO: setting up handlers
	for _, hdl := range s.handlers {
		hdl(s.engine)
	}

	s.Server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.cnf.host, s.cnf.port),
		Handler: s.engine,
	}

	//lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.cnf.host, s.cnf.port))
	//if err != nil {
	//	s.logger.Infof("Listening error: %v", err)
	//	return err
	//}
	//
	//err = s.Serve(lis)
	//
	//if !errors.Is(err, http.ErrServerClosed) {
	//	return err
	//}
	s.logger.Infof("server host %s and port %s and server mode %s", s.cnf.host, s.cnf.port, s.cnf.mode)

	return nil
}

func (s *myHttpServer) Stop(ctx context.Context) error {
	s.logger.Infof("server is shutting down")
	if s.Server != nil {
		_ = s.Server.Shutdown(ctx)
	}
	return nil
}

func (s *myHttpServer) AddHandler(hdl func(e *gin.Engine)) {
	s.handlers = append(s.handlers, hdl)
}