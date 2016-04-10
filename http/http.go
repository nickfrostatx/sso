package http

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nickfrostatx/sso/auth"
	"github.com/nickfrostatx/sso/signer"
	"net/http"
)

const (
	cookieName = "token"
)

type Server struct {
	auth   *auth.Auth
	signer *signer.Signer
}

func NewServer(auth *auth.Auth, signer *signer.Signer) *Server {
	return &Server{
		auth:   auth,
		signer: signer,
	}
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

func (s *Server) NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/login", s.Login)
	return router
}

func (s *Server) Serve() error {
	r := s.NewRouter()
	return http.ListenAndServe("127.0.0.1:8080", r)
}
