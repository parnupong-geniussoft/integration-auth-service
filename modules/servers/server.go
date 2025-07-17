package servers

import (
	"integration-auth-service/configs"
	"integration-auth-service/pkg/utils"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
	Cfg *configs.Configs
	Db  *sqlx.DB
	C   *cache.Cache
}

func NewServer(cfg *configs.Configs, db *sqlx.DB, c *cache.Cache) *Server {
	return &Server{
		App: fiber.New(),
		Cfg: cfg,
		Db:  db,
		C:   c,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	fiberConnURL, err := utils.ConnectionUrlBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.Cfg.App.Host
	port := s.Cfg.App.Port
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
