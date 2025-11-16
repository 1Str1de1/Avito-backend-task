package server

import (
	handler "github.com/1Str1de1/Avito-backend-task/internal/app/handlers"
	"github.com/1Str1de1/Avito-backend-task/internal/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"path"
)

type Server struct {
	router *gin.Engine
	logger *slog.Logger
	config *Config
	db     *model.DB
}

func New(conf *Config) *Server {
	db, err := model.NewDB(conf.env.PostgresUser,
		conf.env.PostgresPassword, conf.env.PostgresHost, conf.env.PostgresName)
	if err != nil {
		panic(err)
	}
	return &Server{
		router: gin.Default(),
		logger: slog.Default(),
		config: conf,
		db:     db,
	}
}

func (s *Server) MustStart(conf *Config) {
	s.configureRouter(conf.env)

	if err := http.ListenAndServe("0.0.0.0:8080", s.router); err != nil {
		panic(err)
	}
}

func (s *Server) configureRouter(env *Env) {
	prApi := s.router.Group(path.Join("api", env.ApiVersion, "pullRequest"))

	prApi.POST("/create", handler.HandlePrCreate(s.db))
	prApi.POST("/merge", handler.HandleMergeRequest(s.db))
	prApi.PUT("/reassign", handler.HandleReassignAuthor(s.db))

	teamApi := s.router.Group(path.Join("api", env.ApiVersion, "team"))

	teamApi.POST("/add", handler.HandleAddTeam(s.db))
	teamApi.GET("/get", handler.HandleGetTeam(s.db))

	usersApi := s.router.Group(path.Join("api", env.ApiVersion, "users"))

	usersApi.GET("/:id", handler.HandleGetUser(s.db))
	usersApi.PUT("/setIsActive", handler.HandleUpdateIsActive(s.db))
}

func (s *Server) Stop() {
	// TODO: Graceful shutdown

}
